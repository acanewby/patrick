package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"github.com/acanewby/patrick/internal/logger"
	"os"
)

func List(cfg Config) {

	setupRun(cfg)

	var fileList []string
	var excludeList []string
	var err error

	// Get a slice of excludable filenames
	if cfg.ExcludeFile != "" {
		if excludeList, err = readTextFileContents(cfg.ExcludeFile); err != nil {
			logger.Errorf(common.ErrorTemplateFileRead, err)
			os.Exit(common.EXIT_CODE_IO_ERROR)
		}
	}

	// Get the list of targeted files as a slice
	if fileList, err = traverseDirectory(cfg.InputDir, excludeList); err != nil {
		for _, f := range fileList {
			logger.Infof(common.LogTemplateFileOpen, f)
		}

	}

	// Produce the output

	fmt.Println(common.UiLabelFilesToProcess)
	singleLine()

	for _, file := range fileList {
		fmt.Println(fmt.Sprintf(common.UiTemplateProcessingFile, file))
	}

}
