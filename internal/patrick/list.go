package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/logger"
	"os"
)

func List(cfg Config) {
	dumpConfig(cfg)

	var fileList []string
	var excludeList []string
	var err error

	if cfg.ExcludeFile != "" {
		if excludeList, err = readFileContents(cfg.ExcludeFile); err != nil {
			logger.Errorf(ErrorTemplateFileRead, err)
			os.Exit(EXIT_CODE_IO_ERROR)
		}
	}

	if fileList, err = traverseDirectory(cfg.InputDir, excludeList); err != nil {
		for _, f := range fileList {
			logger.Infof(LogTemplateFileOpen, f)
		}

	}

	for _, file := range fileList {
		fmt.Println(fmt.Sprintf("Processing: %s", file))
	}

}
