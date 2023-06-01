package patrick

const (
	EXIT_CODE_UNDETERMINED_ERROR int = iota + 1
	EXIT_CODE_INVOCATION_ERROR
	EXIT_CODE_IO_ERROR
)

const (
	FlagInputDir = "inputDir"

	FlagOutputDir = "outputDir"

	FlagExcludeFiles = "excludeFiles"

	ErrorTemplateInvocation            = "invocation error: [%v]"
	ErrorTemplateUndeterminedExecution = "undetermined execution error: [%v]"
	ErrorTemplateFileRead              = "error reading file: [%v]"

	LogTemplateFileRead      = "read file line: [%s]"
	LogTemplateFileOpen      = "opening file: [%s]"
	LogTemplateFileSkip      = "skipping file: [%s]"
	LogTemplateDirectorySkip = "skipping directory: [%s]"
)
