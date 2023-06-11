package common

// traverseWorker performs work
// The parameters can be used to identify work targets, options or behaviors
// An error can be returned
type traverseWorker func(string) error

type Config struct {
	InputDir              string
	OutputDir             string
	OverwriteOutput       bool
	ExcludeFile           string
	LogLevel              string
	ResourceFileDelimiter string
	LanguageConfig        LanguageConfig
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
