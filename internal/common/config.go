package common

var (
	cfg Config
)

func SetConfig(c Config) {
	cfg = c
}

func GetConfig() Config {
	return cfg
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
