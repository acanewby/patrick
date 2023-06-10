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

	cfg := common.GetConfig()

	excludeList := setupRun(cfg)
	var (
		err         error
		isCollision bool
	)

	// Output should not be the same as input, or a child of input
	isCollision, err = common.DirectoryCollision(cfg.InputDir, cfg.OutputDir)
	if err != nil {
		os.Exit(common.EXIT_CODE_UNDETERMINED_ERROR)
	}
	if isCollision {
		os.Exit(common.EXIT_CODE_CONFIGURATION_ERROR)
	}

	// Process the targeted files as a slice
	if err = common.TraverseFilteredDirectoryTree(cfg.InputDir, excludeList, convertFile); err != nil {
		common.LogErrorf(common.ErrorTemplateTraverserExecution, err)
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

	fmt.Println(fmt.Sprintf(common.UiTemplateProcessingFile, inputFilePath))

	// --- setup ------

	// 0. Open input file (IN) for read (err if fail)
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
		common.LogDebugf(common.LogTemplateFileReadLine, line)

		// Trim whitespace and eliminate any trailing comment
		semanticLine, blockCommentBegan, blockCommentEnded := semanticLine(line)

		// 2. Identify package if we don't already know
		if packageName == "" {
			// Can we id the package
			if strings.HasPrefix(semanticLine, cfg.PackageIdentifier) {
				// strip off the identifier
				pkg := strings.TrimSpace(strings.Replace(semanticLine, cfg.PackageIdentifier, "", 1))
				common.LogInfof(common.LogTemplatePackage, pkg)
				// remember the package name
				packageName = pkg
			}
			// 3. Open package resource files (RES) for append
			resourceFilePath := filepath.Join(outputDir, fmt.Sprintf("%s%s", packageName, common.ResourceFileExtension))
			if res, err = common.OpenFileForAppend(resourceFilePath); err != nil {
				return err
			}
			defer common.CloseFile(res)
		}

		priorCodeState := codeState
		codeState = updateCodeState(codeState, semanticLine, blockCommentBegan, blockCommentEnded)

		if shouldParse(priorCodeState, codeState) {
			// 4. Identify string literals
			literals := extractStringLiterals(semanticLine)
			if len(literals) != 0 {
				for _, literal := range literals {
					common.LogDebugf(common.LogTemplateProcessingLiteral, literal)
					if _, err = res.WriteString(fmt.Sprintf("%s\n", literal)); err != nil {
						msg := fmt.Sprintf(common.ErrorTemplateIo, err)
						common.LogErrorf(msg)
						fmt.Println(msg)
						return err
					}
				}
			}

			// 4a. Assign literal token
			// 4b. Substitute token for literal in OUT
			// 4c. Write token and literal to RES
		} else {
			// Pass the line straight through as-is
			if _, err = out.WriteString(line + "\n"); err != nil {
				common.LogErrorf(common.ErrorTemplateIo, err)
				return err
			}
		}

	}

	return nil
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
					examination = examination[0:blockBeginIdx] + examination[blockEndIdx+len(cfg.BlockCommentEndDelimiter):len(examination)]
				} else {
					// Extract the code somewhere in the middle of this line
					examination = examination[blockEndIdx+len(cfg.BlockCommentEndDelimiter) : blockBeginIdx]
				}
			} else {
				// Treat as single-line comment
				examination = examination[0:blockBeginIdx]
				blockCommentBegan = true
			}
		}

		common.LogDebugf(common.LogTemplateFileTrimmedLine, examination)
		singleLineFound, blockBeginFound, blockEndFound, singleLineIdx, blockBeginIdx, blockEndIdx = checkForCommentDelimiters(examination, cfg)
		if blockEndFound {
			if singleLineFound && singleLineIdx > blockEndIdx {
				// Extract the code somewhere in the middle of this line
				examination = examination[blockEndIdx+len(cfg.BlockCommentEndDelimiter) : singleLineIdx]
			} else {
				// Get the code after the block end comment delimiter
				examination = examination[blockEndIdx+len(cfg.BlockCommentEndDelimiter) : len(examination)]
			}
			blockCommentEnded = true
		}

		common.LogDebugf(common.LogTemplateFileTrimmedLine, examination)
		singleLineFound, blockBeginFound, blockEndFound, singleLineIdx, blockBeginIdx, blockEndIdx = checkForCommentDelimiters(examination, cfg)
		if singleLineFound && (!blockBeginFound || singleLineIdx < blockBeginIdx) && (!blockEndFound || singleLineIdx < blockEndIdx) {
			// Single-line comment active
			common.LogDebugf(common.LogTemplateDelimiterPosition, cfg.SingleLineCommentDelimiter, singleLineIdx)
			// Get the code before the single-line comment delimiter
			examination = examination[0:singleLineIdx]
		}

		examination = common.ConsolidateWhitespace(examination)
		common.LogDebugf(common.LogTemplateFileTrimmedLine, examination)

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
	singleLineIdx := strings.Index(examination, cfg.SingleLineCommentDelimiter)
	if singleLineIdx > -1 {
		singleLineFound = true
	}

	blockBeginIdx := strings.Index(examination, cfg.BlockCommentBeginDelimiter)
	if blockBeginIdx > -1 {
		blockBeginFound = true
	}

	blockEndIdx := strings.Index(examination, cfg.BlockCommentEndDelimiter)
	if blockEndIdx > -1 {
		blockEndFound = true
	}
	return singleLineFound, blockBeginFound, blockEndFound, singleLineIdx, blockBeginIdx, blockEndIdx
}

func updateCodeState(currentState codeState, line string, blockCommentBegan bool, blockCommentEnded bool) codeState {

	var newState codeState

	switch {
	case line == "":
		newState = onEmptyLine
	case blockCommentBegan:
		newState = inCommentBlock
	case blockCommentEnded:
		newState = inNormalCode
	case line == constBlockBegin || (currentState == inConstBlock && line != importConstBlockEnd):
		newState = inConstBlock
	case line == importBlockBegin || (currentState == inImportBlock && line != importConstBlockEnd):
		newState = inImportBlock
	case line == importConstBlockEnd && (currentState == inConstBlock || currentState == inImportBlock):
		newState = inNormalCode
	case strings.HasPrefix(line, importKeyword):
		newState = onImportLine
	case strings.HasPrefix(line, constKeyword):
		newState = onConstLine
	default:
		newState = inNormalCode
	}

	// Report new state
	common.LogInfof(newState.String())
	return newState
}

func shouldParse(was codeState, is codeState) bool {

	var parse bool

	/*                   IS
		                     inNormalCode    inCommentBlock    inImportBlock    inConstBlock    onImportLine    onConstLine
		WAS	inNormalCode          T                 T                 F               F               F               F
			inCommentBlock        T                 F                 F               F               F               F
			inImportBlock         T                 T                 F               F               F               F
			inConstBlock          T                 T                 F               F               F               F
	        onImportLine          T                 T                 F               F               F               F
	        onConstLine           T                 T                 F               F               F               F
	*/

	// Skip if the line is empty
	if is == onEmptyLine {
		parse = false
	} else {
		if is == inNormalCode {
			// regardless of what we had before, we now have parseable code
			parse = true
		}

		if is == onImportLine || is == onConstLine {
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
				// even though this line ends with a block comment, we have parseable code
				parse = true
			}
		}
	}

	common.LogDebugf(common.LogTemplateShouldParse, was, is, parse)
	return parse
}
