package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DoubleLineToConsole() {
	fmt.Println(strings.Repeat(DoubleLineChar, ScreenWidth))
}

func SingleLineToConsole() {
	fmt.Println(strings.Repeat(SingleLineChar, ScreenWidth))
}

func ConsolidateWhitespace(line string) string {

	// Top and tail
	result := strings.TrimSpace(line)

	// Replace non-space whitespace
	result = strings.ReplaceAll(result, "\t", "")
	result = strings.ReplaceAll(result, "\n", "")
	result = strings.ReplaceAll(result, "\v", "")
	result = strings.ReplaceAll(result, "\r", "")
	result = strings.ReplaceAll(result, "\f", "")
	result = strings.ReplaceAll(result, "\f", "")
	result = strings.ReplaceAll(result, "\u0085", "")
	result = strings.ReplaceAll(result, "\u00A0", "")

	// Replace duplicate spaces within string
	for {

		// Measure before/after of space consolidation
		startLen := len(result)
		result = strings.ReplaceAll(result, "  ", " ")
		endLen := len(result)

		// We haven't consolidated any more spaces
		if startLen == endLen {
			break
		}

	}

	return result
}

func GetTextFileContents(fileName string) ([]string, error) {

	var err error
	var file *os.File
	var contents []string

	LogDebugf(LogTemplateFileOpen, fileName)

	if file, err = os.Open(fileName); err != nil {
		msg := fmt.Sprintf(ErrorTemplateFileRead, err)
		LogErrorf(msg)
		fmt.Println(msg)
		os.Exit(EXIT_CODE_IO_ERROR)
	}
	defer CloseFile(file)

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		LogDebugf(LogTemplateFileReadLine, line)
		contents = append(contents, line)
	}

	LogDebugf(LogTemplateTextFileContents, contents)
	return contents, nil
}

func IsDirectory(path string) (bool, error) {

	isDir := false

	info, err := os.Stat(path)

	if err != nil {
		LogErrorf(ErrorTemplateIo, err)
		return false, err
	}

	if info.IsDir() {
		isDir = true
	}

	LogDebugf(LogTemplateDirectoryExist, path, isDir)
	return isDir, nil
}

func MkDirP(path string) error {
	LogDebugf(LogTemplateDirCreate, path)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	LogDebugf(LogTemplateDirectoryExist, path, true)
	return nil
}

func TraverseFilteredDirectoryTree(dir string, filter []string, worker traverseWorker) error {

	fmt.Println(UiLabelFilesToProcess)
	SingleLineToConsole()

	var excludeList = make(map[string]string)
	var err error

	for _, f := range filter {
		if f != "" {
			excludeList[f] = f
		}
	}

	if err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {

			// Is this a directory
			if info.IsDir() {
				LogInfof(LogTemplateDirectorySkip, path)
				return nil
			}

			// Should we skip this file
			_, skipFile := excludeList[filepath.Base(path)]
			if skipFile {
				LogInfof(LogTemplateFileSkip, path)
			} else {
				if err = worker(path); err != nil {
					LogErrorf(ErrorTemplateUndeterminedExecution, err)
					return err
				}
			}

			// All good - continue
			return nil

		}); err != nil {
		LogErrorf(ErrorTemplateFileRead, err)
		return err
	}

	SingleLineToConsole()

	return nil
}

// DirectoryCollision verifies that secondaryDir is not the same as, or a child of, primaryDir.
// Returns error if either directory path does not exist or is a file.
func DirectoryCollision(primaryDir string, secondaryDir string) (bool, error) {

	LogInfof(LogTemplateCheckDirectoryCollision, LogPrimaryDir, primaryDir, LogSecondaryDir, secondaryDir)

	var (
		err          error
		absPrimary   string
		absSecondary string
	)
	checklist := []string{primaryDir, secondaryDir}

	// Are they both directories
	for _, dir := range checklist {
		isDir, err := IsDirectory(dir)
		if err != nil {
			LogErrorf(ErrorTemplateIo, err)
			return false, err
		}
		if !isDir {
			msg := fmt.Sprintf(LogTemplateDirectoryExist, dir, isDir)
			LogDebugf(msg)
			return false, errors.New(msg)
		}
	}

	// Get the full paths
	absPrimary, err = filepath.Abs(primaryDir)
	if err != nil {
		LogErrorf(ErrorTemplateIo, err)
		return false, err
	}

	absSecondary, err = filepath.Abs(secondaryDir)
	if err != nil {
		LogErrorf(ErrorTemplateIo, err)
		return false, err
	}
	LogDebugf(LogPrimaryDir+": [%s]", absPrimary)
	LogDebugf(LogSecondaryDir+": [%s]", absSecondary)

	// Are they the same
	if absPrimary == absSecondary {
		LogInfof(LogTemplatePathsMatch, absPrimary)
		fmt.Println(fmt.Sprintf(UiTemplateDirCollision, primaryDir, secondaryDir))
		return true, nil
	}

	// Is secondary a child of primary?
	if strings.HasPrefix(absSecondary, absPrimary) {
		LogInfof(LogTemplatePathCollision, absSecondary, absPrimary)
		fmt.Println(fmt.Sprintf(UiTemplateDirCollision, primaryDir, secondaryDir))
		return true, nil
	}

	// Is secondary an ancestor of primary?
	if strings.HasPrefix(absPrimary, absSecondary) {
		fmt.Println(fmt.Sprintf(UiTemplateDirCollision, primaryDir, secondaryDir))
		LogInfof(LogTemplatePathCollision, absSecondary, absPrimary)
		return true, nil
	}

	// If we are here, there is no identifiable collision
	LogInfof(LogNoPathCollision)
	return false, nil
}

func OpenFileForRead(path string) (*os.File, error) {
	var (
		in  *os.File
		err error
	)
	LogDebugf(LogTemplateFileOpen, path)
	if in, err = os.Open(path); err != nil {
		LogErrorf(ErrorTemplateFileRead, err)
		return nil, err
	}
	return in, nil
}

func OpenFileForOverwrite(path string) (*os.File, error) {
	var (
		out *os.File
		err error
	)
	LogDebugf(LogTemplateFileOpen, path)
	if out, err = os.Create(path); err != nil {
		LogErrorf(ErrorTemplateFileWrite, err)
		return nil, err
	}

	return out, nil
}

func OpenFileForAppend(path string) (*os.File, error) {
	var (
		out *os.File
		err error
	)
	LogDebugf(LogTemplateFileOpen, path)
	if out, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		LogErrorf(ErrorTemplateFileWrite, err)
		return nil, err
	}

	return out, nil
}

func CloseFile(f *os.File) {
	LogDebugf(LogTemplateFileClose, f.Name())
	err := f.Close()
	if err != nil {
		msg := fmt.Sprintf(ErrorTemplateIo, err)
		LogErrorf(msg)
		fmt.Println(msg)
		os.Exit(EXIT_CODE_IO_ERROR)
	}
}
