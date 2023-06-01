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
