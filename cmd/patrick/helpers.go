package patrick

import (
	"github.com/acanewby/patrick/internal/common"
)

var (
	goLanguage = common.LanguageConfig{
		PackageIdentifier:          "package",
		StringDelimiter:            "\"",
		SingleLineCommentDelimiter: "//",
		BlockCommentBeginDelimiter: "/*",
		BlockCommentEndDelimiter:   "*/",
		ImportKeyword:              "import",
		ImportBlockBegin:           "import (",
		ImportBlockEnd:             ")",
		ConstKeyword:               "const",
		ConstBlockBegin:            "const (",
		ConstBlockEnd:              ")",
	}
)

// Right now, there is only a hard-coded language config for Go, but this could be extended
func constructConfig() common.Config {

	return common.Config{
		InputDir:              inputDir,
		OutputDir:             outputDir,
		ExcludeFile:           excludedNamesFile,
		LogLevel:              logLevel,
		ResourceFileDelimiter: resourceFileDelimiter,
		LanguageConfig:        goLanguage,
	}
}
