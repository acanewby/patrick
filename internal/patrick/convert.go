package patrick

import (
	"bufio"
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"os"
	"path/filepath"
	"strings"
)

func Convert() {

	var (
		err             error
		isCollision     bool
		outputDirExists bool
	)

	cfg := common.GetConfig()
	excludeList := setupRun(cfg)

	// Validate resource function template
	if !strings.Contains(cfg.ResourceFunctionTemplate, common.ResourceFunctionTemplateSubstitutionToken) {
		msg := fmt.Sprintf(common.ErrorTemplateResourceFunctionMissingSubToken, common.ResourceFunctionTemplateSubstitutionToken, cfg.ResourceFunctionTemplate)
		common.LogErrorf(msg)
		fmt.Println(msg)
		os.Exit(common.EXIT_CODE_CONFIGURATION_ERROR)
	}

	// Output should not be the same as input, or a child of input
	isCollision, err = common.DirectoryCollision(cfg.InputDir, cfg.OutputDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(common.EXIT_CODE_UNDETERMINED_ERROR)
	}
	if isCollision {
		fmt.Println(fmt.Sprintf(common.LogTemplatePathCollision, cfg.InputDir, cfg.OutputDir))
		os.Exit(common.EXIT_CODE_CONFIGURATION_ERROR)
	}

	// Prepare output directory
	outputDirExists, err = common.IsDirectory(cfg.OutputDir)
	// bale because we barfed trying to stat the path
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(common.EXIT_CODE_UNDETERMINED_ERROR)
	}
	if outputDirExists {
		if !cfg.OverwriteOutput {
			// it exists and we can't overwrite it
			fmt.Println(fmt.Sprintf(common.ErrorTemplateOutputDirAlreadyExists, cfg.OutputDir))
			os.Exit(common.EXIT_CODE_CONFIGURATION_ERROR)
		} else {
			// it exists but we can kill it
			if err = common.RmDirP(cfg.OutputDir); err != nil {
				msg := fmt.Sprintf(common.ErrorTemplateIo, err)
				common.LogErrorf(msg)
				fmt.Println(msg)
				os.Exit(common.EXIT_CODE_IO_ERROR)
			}
		}
	}
	// At this point, the outputdir does not exist, so we must make it
	if err = common.MkDirP(cfg.OutputDir); err != nil {
		msg := fmt.Sprintf(common.ErrorTemplateIo, err)
		common.LogErrorf(msg)
		fmt.Println(msg)
		os.Exit(common.EXIT_CODE_IO_ERROR)
	}

	// Process the targeted files as a slice
	if err = common.TraverseFilteredDirectoryTree(cfg.InputDir, excludeList, convertFile); err != nil {
		msg := fmt.Sprintf(common.ErrorTemplateTraverserExecution, err)
		common.LogErrorf(msg)
		fmt.Println(msg)
		os.Exit(common.EXIT_CODER_TRAVERSER_EXECUTION)
	}

}

/*
convertFile implements traverseWorker.
It reads a given source file, identifies string literals.
It outputs two files:
  - a converted version of the input file, with system-generated tokens in place of the identified literals
  - a resource file with token:literal mappings
*/
func convertFile(inputFilePath string) error {

	var (
		err         error
		in          *os.File
		out         *os.File
		res         *os.File
		packageName string
		codeState   = inNormalCode
	)

	cfg := common.GetConfig()
	inv := common.GetResourceInventory()

	fmt.Println(fmt.Sprintf(common.UiTemplateProcessingFile, inputFilePath))

	// --- setup ------

	// 0a. Open input file (IN) for read (err if fail)
	if in, err = common.OpenFileForRead(inputFilePath); err != nil {
		return err
	}
	defer common.CloseFile(in)

	// 1. Open output file (OUT) for write (err if already exists?)
	outputFilePath, outputDir := determineOutputDestinations(inputFilePath, cfg)

	if err = common.MkDirP(outputDir); err != nil {
		common.LogErrorf(common.ErrorTemplateIo, err)
		return err
	}

	if out, err = common.OpenFileForOverwrite(outputFilePath); err != nil {
		return err
	}
	defer common.CloseFile(out)

	// --- process IN line by line ------
	// inBlockComment := false

	fileScanner := bufio.NewScanner(in)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		common.LogInfof(common.LogTemplateFileReadLine, line)

		// Trim whitespace and eliminate any trailing comment
		semanticLine, blockCommentBegan, blockCommentEnded := semanticLine(line)

		// 2. Identify package if we don't already know
		if packageName == "" {
			// Can we id the package
			packageName = identifyPackage(semanticLine, cfg)
			// 3. Open package resource files (RES) for append
			resourceFilePath := filepath.Join(outputDir, fmt.Sprintf("%s%s", packageName, common.ResourceFileExtension))
			if res, err = common.OpenFileForAppend(resourceFilePath); err != nil {
				return err
			}
			defer common.CloseFile(res)
		}

		priorCodeState := codeState
		codeState = updateCodeState(codeState, semanticLine, blockCommentBegan, blockCommentEnded)

		parse := shouldParse(priorCodeState, codeState)
		if parse {
			// 4. Identify string literals
			literals := extractStringLiterals(semanticLine)
			if len(literals) != 0 {
				for _, resource := range literals {

					if resource == common.StringRepresentingEmptyString {
						common.LogDebugf(common.LogTemplateIgnoringEmptyLiteral, resource)
					} else {

						common.LogDebugf(common.LogTemplateProcessingLiteral, resource)

						// get the resource index for this token
						idx, isNew := inv.GetIndexForResource(packageName, resource)
						resourceToken := inv.ResourceToken(idx)

						// If it's a new one, also write to the resource file
						if isNew {
							entry := resourceFileEntry(resourceToken, resource)
							if _, err = res.WriteString(entry + "\n"); err != nil {
								msg := fmt.Sprintf(common.ErrorTemplateIo, err)
								common.LogErrorf(msg)
								fmt.Println(msg)
								return err
							}
						}

						// Update the code line
						resourceFunctionCall := inv.GetResourceFunctionCall(resourceToken)
						line = strings.Replace(line, resource, resourceFunctionCall, 1)
					}

				}
			}
		}

		// Write the line to output
		if _, err = out.WriteString(line + "\n"); err != nil {
			common.LogErrorf(common.ErrorTemplateIo, err)
			return err
		}
		common.LogDebugf(common.LogTemplateFileWroteLine, line)

	}

	return nil
}

func resourceFileEntry(resourceToken string, resource string) string {

	cfg := common.GetConfig()

	entry := fmt.Sprintf("%s%s%s%s%s", cfg.ResourceFilePrefix, resourceToken, cfg.ResourceFileDelimiter, resource, cfg.ResourceFileSuffix)

	common.LogDebugf(common.LogTemplateResourceEntry, entry)

	return entry
}

func identifyPackage(semanticLine string, cfg common.Config) string {

	packageName := ""

	if strings.HasPrefix(semanticLine, cfg.LanguageConfig.PackageIdentifier) {
		// strip off the identifier
		pkg := strings.TrimSpace(strings.Replace(semanticLine, cfg.LanguageConfig.PackageIdentifier, "", 1))
		common.LogInfof(common.LogTemplatePackage, pkg)
		// remember the package name
		packageName = pkg
	}
	return packageName
}

func determineOutputDestinations(path string, cfg common.Config) (string, string) {
	projectFilename := strings.Replace(path, cfg.InputDir, "", 1)
	outputPath := filepath.Join(cfg.OutputDir, projectFilename)
	common.LogInfof(common.LogTemplateFileOutput, outputPath)
	outputDir := strings.Replace(outputPath, filepath.Base(outputPath), "", 1)
	return outputPath, outputDir
}

func semanticLine(line string) (string, bool, bool) {

	var (
		singleLineFound   bool
		blockBeginFound   bool
		blockEndFound     bool
		singleLineIdx     int
		blockBeginIdx     int
		blockEndIdx       int
		blockCommentBegan bool
		blockCommentEnded bool
	)

	cfg := common.GetConfig()

	// Strip leading/trailing whitespace
	examination := strings.TrimSpace(line)

	// Loop until we haven't trimmed anything
	for {

		// Remember starting length
		startLen := len(examination)

		common.LogDebugf(common.LogTemplateFileTrimmedLine, examination)
		singleLineFound, blockBeginFound, blockEndFound, singleLineIdx, blockBeginIdx, blockEndIdx = checkForCommentDelimiters(examination, cfg)
		if blockBeginFound {
			if blockEndFound {
				if blockEndIdx > blockBeginIdx {
					// Remove the block comment somewhere in the middle of this line
					examination = examination[0:blockBeginIdx] + examination[blockEndIdx+len(cfg.LanguageConfig.BlockCommentEndDelimiter):len(examination)]
				} else {
					// Extract the code somewhere in the middle of this line
					examination = examination[blockEndIdx+len(cfg.LanguageConfig.BlockCommentEndDelimiter) : blockBeginIdx]
				}
			} else {
				// Treat as single-line comment and exit loop
				examination = examination[0:blockBeginIdx]
				blockCommentBegan = true
				// If the line is now empty, we can exit the loop
				if examination == "" {
					break
				}
			}
		}

		common.LogDebugf(common.LogTemplateFileTrimmedLine, examination)
		singleLineFound, blockBeginFound, blockEndFound, singleLineIdx, blockBeginIdx, blockEndIdx = checkForCommentDelimiters(examination, cfg)
		if blockEndFound {
			if singleLineFound && singleLineIdx > blockEndIdx {
				// Extract the code somewhere in the middle of this line
				examination = examination[blockEndIdx+len(cfg.LanguageConfig.BlockCommentEndDelimiter) : singleLineIdx]
			} else {
				// Get the code after the block end comment delimiter
				examination = examination[blockEndIdx+len(cfg.LanguageConfig.BlockCommentEndDelimiter) : len(examination)]
			}
			blockCommentEnded = true
			// If the line is now empty, we can exit the loop
			if examination == "" {
				break
			}
		}

		common.LogDebugf(common.LogTemplateFileTrimmedLine, examination)
		singleLineFound, blockBeginFound, blockEndFound, singleLineIdx, blockBeginIdx, blockEndIdx = checkForCommentDelimiters(examination, cfg)
		if singleLineFound && (!blockBeginFound || singleLineIdx < blockBeginIdx) && (!blockEndFound || singleLineIdx < blockEndIdx) {
			// Single-line comment active
			common.LogDebugf(common.LogTemplateDelimiterPosition, cfg.LanguageConfig.SingleLineCommentDelimiter, singleLineIdx)
			// Get the code before the single-line comment delimiter
			examination = examination[0:singleLineIdx]
		}

		examination = common.ConsolidateWhitespace(examination)
		common.LogInfof(common.LogTemplateFileTrimmedLine, examination)

		// Did we trim anything?
		if startLen == len(examination) {
			break
		}
	}

	return examination, blockCommentBegan, blockCommentEnded
}

func checkForCommentDelimiters(examination string, cfg common.Config) (bool, bool, bool, int, int, int) {
	// comment state possibilities
	singleLineFound := false
	blockBeginFound := false
	blockEndFound := false

	// Look for comment delimiters
	singleLineIdx := strings.Index(examination, cfg.LanguageConfig.SingleLineCommentDelimiter)
	if singleLineIdx > -1 {
		singleLineFound = true
	}

	blockBeginIdx := strings.Index(examination, cfg.LanguageConfig.BlockCommentBeginDelimiter)
	if blockBeginIdx > -1 {
		blockBeginFound = true
	}

	blockEndIdx := strings.Index(examination, cfg.LanguageConfig.BlockCommentEndDelimiter)
	if blockEndIdx > -1 {
		blockEndFound = true
	}
	return singleLineFound, blockBeginFound, blockEndFound, singleLineIdx, blockBeginIdx, blockEndIdx
}

func updateCodeState(prevState codeState, line string, blockCommentBegan bool, blockCommentEnded bool) codeState {

	var newState codeState
	cfg := common.GetConfig()

	switch {
	// We must evaluate block comments first, since they will be semantically resolved to an empty line if there is no actual code
	case blockCommentBegan || (prevState == inCommentBlock && !blockCommentEnded):
		newState = inCommentBlock
	case blockCommentEnded:
		newState = inNormalCode
	case line == cfg.LanguageConfig.ConstBlockBegin || (prevState == inConstBlock && line != cfg.LanguageConfig.ConstBlockEnd):
		newState = inConstBlock
	case line == cfg.LanguageConfig.ImportBlockBegin || (prevState == inImportBlock && line != cfg.LanguageConfig.ImportBlockEnd):
		newState = inImportBlock
	case line == cfg.LanguageConfig.ImportBlockEnd && prevState == inImportBlock:
		newState = inNormalCode
	case line == cfg.LanguageConfig.ConstBlockEnd && prevState == inConstBlock:
		newState = inNormalCode
	case strings.HasPrefix(line, cfg.LanguageConfig.ImportKeyword):
		newState = onImportLine
	case strings.HasPrefix(line, cfg.LanguageConfig.ConstKeyword):
		newState = onConstLine
	case strings.HasPrefix(line, cfg.LanguageConfig.PackageIdentifier):
		newState = onPackageLine
		// The empty line state should only be returned when we are otherwise in normal code, so we can avoid parsing
	case line == "":
		newState = onEmptyLine
	default:
		newState = inNormalCode
	}

	// Report new state
	common.LogDebugf(newState.String())
	return newState
}

func shouldParse(was codeState, is codeState) bool {

	var parse bool

	/*                   IS
			                     inNormalCode    inCommentBlock    inImportBlock    inConstBlock    onImportLine    onConstLine    onPackageLine
			WAS	inNormalCode          T                 T                 F               F               F               F               F
				inCommentBlock        T                 F                 F               F               F               F               F
				inImportBlock         T                 T                 F               F               F               F               F
				inConstBlock          T                 T                 F               F               F               F               F
		        onImportLine          T                 T                 F               F               F               F               F
		        onConstLine           T                 T                 F               F               F               F               F
	            onPackageLine         T                 T                 F               F               F               F               F
	*/

	// Skip if the line is empty
	if is == onEmptyLine {
		parse = false
	} else {
		if is == inNormalCode {
			// regardless of what we had before, we now have parseable code
			parse = true
		}

		if is == onImportLine || is == onConstLine || is == onPackageLine {
			// regardless of what we had before, we are now in part of the code we should not touch
			parse = false
		}

		if is == inImportBlock || is == inConstBlock {
			// regardless of what we had before, we are now in part of the code we should not touch
			parse = false
		}

		if is == inCommentBlock {
			if was == inCommentBlock {
				// we are still in the comment block we were last iteration
				parse = false
			} else {
				// even though this line ends with a block comment, we may have parseable code
				parse = true
			}
		}
	}

	common.LogInfof(common.LogTemplateShouldParse, was, is, parse)
	return parse
}
