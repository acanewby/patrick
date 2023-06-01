package patrick

import (
	"bufio"
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"github.com/acanewby/patrick/internal/logger"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	InputDir    string
	OutputDir   string
	ExcludeFile string
	LogLevel    string
}

func setupRun(cfg Config) {
	if cfg.LogLevel != "" {
		logger.SetLogLevel(cfg.LogLevel)
	}

	dumpConfig(cfg)
}

func dumpConfig(cfg Config) {
	doubleLine()

	logger.Infof(common.LogTemplateConfig, cfg)

	fmt.Println(fmt.Sprintf(common.UiTemplateInputDir, cfg.InputDir))
	fmt.Println(fmt.Sprintf(common.UiTemplateOutputDir, cfg.OutputDir))
	fmt.Println(fmt.Sprintf(common.UiTemplateExcludesFile, cfg.ExcludeFile))
	fmt.Println(fmt.Sprintf(common.UiTemplateLogLevel, cfg.LogLevel))

	doubleLine()
}

func doubleLine() {
	fmt.Println(strings.Repeat(common.DoubleLineChar, common.ScreenWidth))
}

func singleLine() {
	fmt.Println(strings.Repeat(common.SingleLineChar, common.ScreenWidth))
}

func readTextFileContents(fileName string) ([]string, error) {

	var err error
	var file *os.File
	var contents []string

	logger.Infof(common.LogTemplateFileOpen, fileName)

	if file, err = os.Open(fileName); err != nil {
		logger.Errorf(common.ErrorTemplateFileRead, err)
		os.Exit(common.EXIT_CODE_IO_ERROR)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		logger.Debugf(common.LogTemplateFileRead, line)
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
				logger.Infof(common.LogTemplateDirectorySkip, path)
				return nil
			}

			// Should we skip this file
			_, skipFile := excludeList[filepath.Base(path)]
			if skipFile {
				logger.Infof(common.LogTemplateFileSkip, path)
			} else {
				filelist = append(filelist, path)
			}

			// All good - continue
			return nil

		}); err != nil {
		logger.Errorf(common.ErrorTemplateFileRead, err)
		os.Exit(common.EXIT_CODE_IO_ERROR)
	}

	return filelist, nil
}
