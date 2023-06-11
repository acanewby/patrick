package patrick

import (
	"github.com/acanewby/patrick/internal/common"
)

// Right now, there is only a hard-coded language config for Go, but this could be extended
func constructConfig() common.Config {

	return common.Config{
		InputDir:              inputDir,
		OutputDir:             outputDir,
		OverwriteOutput:       overwriteOutput,
		ExcludeFile:           excludedNamesFile,
		LogLevel:              logLevel,
		ResourceFileDelimiter: resourceFileDelimiter,
		LanguageConfig:        common.GoLanguageConfig(),
	}
}
