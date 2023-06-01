package common

const (
	EXIT_CODE_UNDETERMINED_ERROR int = iota + 1
	EXIT_CODE_INVOCATION_ERROR
	EXIT_CODE_IO_ERROR
	EXIT_CODE_CONFIGURATION_ERROR
)

const (
	ScreenWidth    = 80
	DoubleLineChar = "="
	SingleLineChar = "-"

	FlagLogLevel     = "logLevel"
	FlagInputDir     = "inputDir"
	FlagOutputDir    = "outputDir"
	FlagExcludeFiles = "excludeFiles"

	UiLabelFilesToProcess    = "Files to process"
	UiTemplateProcessingFile = "Processing file: %s"
	UiTemplateInputDir       = "Input directory  : %s"
	UiTemplateOutputDir      = "Output directory : %s"
	UiTemplateExcludesFile   = "Exclude files    : %s"
	UiTemplateLogLevel       = "Log level        : %s"

	ErrorTemplateInvocation            = "invocation error: [%v]"
	ErrorTemplateUndeterminedExecution = "undetermined execution error: [%v]"
	ErrorTemplateFileRead              = "error reading file: [%v]"
	ErrorTemplateParseError            = "parse error: [%+v]"

	LogTemplateFileRead        = "read file line: [%s]"
	LogTemplateFileOpen        = "opening file: [%s]"
	LogTemplateFileSkip        = "skipping file: [%s]"
	LogTemplateDirectorySkip   = "skipping directory: [%s]"
	LogTemplateConfig          = "config: [%+v]"
	LogTemplateSettingLogLevel = "setting log level: [%v]"
	LogTemplateSetLogLevel     = "set log level: [%v]"
)
