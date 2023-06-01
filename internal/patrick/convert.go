package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"os"
)

func Convert(cfg Config) {

	var (
		err error
	)

	setupRun(cfg)

	// Output should not be the same as input, or a child of input
	isCollision, err := common.DirectoryCollision(cfg.InputDir, cfg.OutputDir)
	if err != nil {
		common.LogErrorf(common.ErrorTemplateIo, err)
		os.Exit(common.EXIT_CODE_IO_ERROR)
	}
	if isCollision {
		msg := fmt.Sprintf(common.LogTemplatePathCollision, cfg.OutputDir, cfg.InputDir)
		fmt.Println(msg)
		common.LogErrorf(msg)
		os.Exit(common.EXIT_CODE_CONFIGURATION_ERROR)
	}

}
