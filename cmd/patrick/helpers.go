package patrick

import (
	"github.com/acanewby/patrick/internal/common"
)

func constructConfig() common.Config {
	return common.Config{
		InputDir:                   inputDir,
		OutputDir:                  outputDir,
		ExcludeFile:                excludedNamesFile,
		LogLevel:                   logLevel,
		PackageIdentifier:          packageIdentifier,
		ResourceFileDelimiter:      resourceFileDelimiter,
		StringDelimiter:            stringDelimiter,
		SingleLineCommentDelimiter: singleCommentDelimiter,
		BlockCommentBeginDelimiter: blockCommentBeginDelimiter,
		BlockCommentEndDelimiter:   blockCommentEndDelimiter,
	}
}
