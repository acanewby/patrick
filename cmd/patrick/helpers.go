package patrick

import "github.com/acanewby/patrick/internal/patrick"

func constructConfig() patrick.Config {
	return patrick.Config{
		InputDir:    inputDir,
		OutputDir:   outputDir,
		ExcludeFile: excludedNamesFile,
		LogLevel:    logLevel,
	}
}
