package patrick

import (
	"github.com/acanewby/patrick/internal/common"
	"strings"
)

func constructConfig() common.Config {
	return common.Config{
		InputDir:                   inputDir,
		OutputDir:                  outputDir,
		ExcludeFile:                excludedNamesFile,
		LogLevel:                   logLevel,
		PackageIdentifier:          packageIdentifier,
		ResourceFileDelimiter:      resourceFileDelimiter,
		StringDelimiters:           strings.Split(stringDelimiters, ","),
		SingleLineCommentDelimiter: singleCommentDelimiter,
		BlockCommentBeginDelimiter: blockCommentBeginDelimiter,
		BlockCommentEndDelimiter:   blockCommentEndDelimiter,
	}
}
