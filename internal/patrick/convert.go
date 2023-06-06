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
	)

	cfg := common.GetConfig()

	fmt.Println(fmt.Sprintf(common.UiTemplateProcessingFile, inputFilePath))

	// --- setup ------

	// 0. Open input file (IN) for read (err if fail)
	if in, err = common.OpenFileForRead(inputFilePath); err != nil {
		return err
	}
	defer func(in *os.File) {
		common.LogDebugf(common.LogTemplateFileClose, in.Name())
		err := in.Close()
		if err != nil {
			msg := fmt.Sprintf(common.ErrorTemplateIo, err)
			common.LogErrorf(msg)
			fmt.Println(msg)
			os.Exit(common.EXIT_CODE_IO_ERROR)
		}
	}(in)

	// 1. Open output file (OUT) for write (err if already exists?)
	outputFilePath, outputDir := determineOutputDestinations(inputFilePath, cfg)

	if err = common.MkDirP(outputDir); err != nil {
		common.LogErrorf(common.ErrorTemplateIo, err)
		return err
	}

	if out, err = common.OpenFileForOverwrite(outputFilePath); err != nil {
		return err
	}
	defer func(out *os.File) {
		common.LogDebugf(common.LogTemplateFileClose, out.Name())
		err := out.Close()
		if err != nil {
			msg := fmt.Sprintf(common.ErrorTemplateIo, err)
			common.LogErrorf(msg)
			fmt.Println(msg)
			os.Exit(common.EXIT_CODE_IO_ERROR)
		}
	}(out)

	// --- process IN line by line ------
	// inBlockComment := false

	fileScanner := bufio.NewScanner(in)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		common.LogDebugf(common.LogTemplateFileReadLine, line)

		// Trim whitespace and eliminate any trailing comment
		semanticLine := semanticLine(line)

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
		}

		// 3. Open package resource files (RES) for append
		resourceFilePath := filepath.Join(outputDir, fmt.Sprintf("%s%s", packageName, common.ResourceFileExtension))
		if res, err = common.OpenFileForAppend(resourceFilePath); err != nil {
			return err
		}
		defer func(res *os.File) {
			common.LogDebugf(common.LogTemplateFileClose, res.Name())
			err := res.Close()
			if err != nil {
				msg := fmt.Sprintf(common.ErrorTemplateIo, err)
				common.LogErrorf(msg)
				fmt.Println(msg)
				os.Exit(common.EXIT_CODE_IO_ERROR)
			}
		}(res)
		res.WriteString("hello\n")

		// 4. Identify string literals

		// We must throw an error if we don't know the package name and have the pkg resource file open

		// 4a. Assign literal token
		// 4b. Substitute token for literal in OUT
		// 4c. Write token and literal to RES

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

func semanticLine(line string) string {

	var (
		singleLineFound bool
		blockBeginFound bool
		blockEndFound   bool
		singleLineIdx   int
		blockBeginIdx   int
		blockEndIdx     int
	)

	cfg := common.GetConfig()

	// Strip leading/trailing whitespace
	examination := strings.TrimSpace(line)

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
	return examination
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
