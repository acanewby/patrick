package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/common"
)

type Config struct {
	InputDir    string
	OutputDir   string
	ExcludeFile string
	LogLevel    string
}

func setupRun(cfg Config) {
	if cfg.LogLevel != "" {
		common.SetLogLevel(cfg.LogLevel)
	}

	dumpConfig(cfg)
}

func dumpConfig(cfg Config) {
	common.DoubleLineToConsole()

	common.LogInfof(common.LogTemplateConfig, cfg)

	fmt.Println(fmt.Sprintf(common.UiTemplateInputDir, cfg.InputDir))
	fmt.Println(fmt.Sprintf(common.UiTemplateOutputDir, cfg.OutputDir))
	fmt.Println(fmt.Sprintf(common.UiTemplateExcludesFile, cfg.ExcludeFile))
	fmt.Println(fmt.Sprintf(common.UiTemplateLogLevel, cfg.LogLevel))

	common.DoubleLineToConsole()
}
