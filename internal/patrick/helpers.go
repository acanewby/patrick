package patrick

import (
	"bufio"
	"fmt"
	"github.com/acanewby/patrick/internal/logger"
	"os"
	"path/filepath"
)

type Config struct {
	InputDir    string
	OutputDir   string
	ExcludeFile string
}

func dumpConfig(cfg Config) {
	doubleLine()

	logger.Infof("config: %+v", cfg)

	fmt.Println(fmt.Sprintf("Input directory : %s", cfg.InputDir))
	fmt.Println(fmt.Sprintf("Output directory: %s", cfg.OutputDir))
	fmt.Println(fmt.Sprintf("Exclude files   : %s", cfg.ExcludeFile))

	doubleLine()
}

func doubleLine() {
	fmt.Println("================================================================================")
}

func singleLine() {
	fmt.Println("--------------------------------------------------------------------------------")
}

func readFileContents(fileName string) ([]string, error) {

	var err error
	var file *os.File
	var contents []string

	logger.Infof(LogTemplateFileOpen, fileName)

	if file, err = os.Open(fileName); err != nil {
		logger.Errorf(ErrorTemplateFileRead, err)
		os.Exit(EXIT_CODE_IO_ERROR)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		logger.Debugf(LogTemplateFileRead, line)
		contents = append(contents, line)
	}

	return contents, nil
}

func traverseDirectory(dir string, filter []string) ([]string, error) {

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
				logger.Infof(LogTemplateDirectorySkip, path)
				return nil
			}

			// Should we skip this file
			_, skipFile := excludeList[filepath.Base(path)]
			if skipFile {
				logger.Infof(LogTemplateFileSkip, path)
			} else {
				filelist = append(filelist, path)
			}

			// All good - continue
			return nil

		}); err != nil {
		logger.Errorf(ErrorTemplateFileRead, err)
		os.Exit(EXIT_CODE_IO_ERROR)
	}

	return filelist, nil
}
