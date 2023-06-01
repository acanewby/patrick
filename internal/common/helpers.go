package common

import (
	"bufio"
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

func ReadTextFileContents(fileName string) ([]string, error) {

	var err error
	var file *os.File
	var contents []string

	LogInfof(LogTemplateFileOpen, fileName)

	if file, err = os.Open(fileName); err != nil {
		LogErrorf(ErrorTemplateFileRead, err)
		os.Exit(EXIT_CODE_IO_ERROR)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		LogDebugf(LogTemplateFileRead, line)
		contents = append(contents, line)
	}

	return contents, nil
}

func IsDirectory(path string) (bool, error) {

	isDir := false

	info, err := os.Stat(path)

	if err != nil {
		LogErrorf(ErrorTemplateIo, err)
		os.Exit(EXIT_CODE_IO_ERROR)
	}

	if info.IsDir() {
		isDir = true
	}

	LogDebugf(LogTemplateDirectoryExist, path, isDir)
	return isDir, nil
}

func FilteredDirectoryTreeFiles(dir string, filter []string) ([]string, error) {

	var filelist []string
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
				filelist = append(filelist, path)
			}

			// All good - continue
			return nil

		}); err != nil {
		LogErrorf(ErrorTemplateFileRead, err)
		os.Exit(EXIT_CODE_IO_ERROR)
	}

	return filelist, nil
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
			os.Exit(EXIT_CODE_IO_ERROR)
		}
		if !isDir {
			LogDebugf(LogTemplateDirectoryExist, dir, isDir)
			os.Exit(EXIT_CODE_CONFIGURATION_ERROR)
		}
	}

	// Get the full paths
	absPrimary, err = filepath.Abs(primaryDir)
	if err != nil {
		LogErrorf(ErrorTemplateIo, err)
		os.Exit(EXIT_CODE_IO_ERROR)
	}

	absSecondary, err = filepath.Abs(secondaryDir)
	if err != nil {
		LogErrorf(ErrorTemplateIo, err)
		os.Exit(EXIT_CODE_IO_ERROR)
	}
	LogDebugf(LogPrimaryDir+": [%s]", absPrimary)
	LogDebugf(LogSecondaryDir+": [%s]", absSecondary)

	// Are they the same
	if absPrimary == absSecondary {
		LogInfof(LogTemplatePathsMatch, absPrimary)
		return true, nil
	}

	// Is secondary a child of primary?
	if strings.HasPrefix(absSecondary, absPrimary) {
		LogInfof(LogTemplatePathCollision, absSecondary, absPrimary)
		return true, nil
	}

	// If we are here, there is no identifiable collision
	LogInfof(LogNoPathCollision)
	return false, nil
}
