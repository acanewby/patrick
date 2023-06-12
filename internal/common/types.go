package common

// traverseWorker performs work
// The parameters can be used to identify work targets, options or behaviors
// An error can be returned
type traverseWorker func(string) error

type Config struct {
	InputDir                 string
	OutputDir                string
	OverwriteOutput          bool
	ExcludeFile              string
	LogLevel                 string
	ResourceFileDelimiter    string
	ResourceIndexStart       uint64
	ResourceIndexZeroPad     uint8
	ResourceTokenPrefix      string
	ResourceFunctionTemplate string
	LanguageConfig           LanguageConfig
}

type LanguageConfig struct {
	PackageIdentifier          string
	StringDelimiter            string
	SingleLineCommentDelimiter string
	BlockCommentBeginDelimiter string
	BlockCommentEndDelimiter   string
	ImportKeyword              string
	ImportBlockBegin           string
	ImportBlockEnd             string
	ConstKeyword               string
	ConstBlockBegin            string
	ConstBlockEnd              string
}

// GoLanguageConfig returns a LanguageConfig specific to Go(lang)
func GoLanguageConfig() LanguageConfig {
	return LanguageConfig{
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
}
