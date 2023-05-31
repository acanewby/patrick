package patrick

const (
	EXIT_CODE_UNDETERMINED_ERROR int = iota + 1
	EXIT_CODE_INVOCATION_ERROR
)

const (
	FlagInputDir = "inputDir"

	FlagOutputDir = "outputDir"

	FlagExcludeFiles = "excludeFiles"

	ErrorTemplateInvocation            = "invocation error: [%v]"
	ErrorTemplateUndeterminedExecution = "undetermined execution error: [%v]"
)
