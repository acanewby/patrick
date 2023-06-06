package common

// traverseWorker performs work
// The parameters can be used to identify work targets, options or behaviors
// An error can be returned
type traverseWorker func(string) error

type Config struct {
	InputDir                   string
	OutputDir                  string
	ExcludeFile                string
	LogLevel                   string
	PackageIdentifier          string
	ResourceFileDelimiter      string
	StringDelimiter            string
	SingleLineCommentDelimiter string
	BlockCommentBeginDelimiter string
	BlockCommentEndDelimiter   string
}
