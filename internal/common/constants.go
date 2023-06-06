package common

const (
	EXIT_CODE_UNDETERMINED_ERROR int = iota + 1
	EXIT_CODE_INVOCATION_ERROR
	EXIT_CODE_IO_ERROR
	EXIT_CODE_CONFIGURATION_ERROR
	EXIT_CODER_TRAVERSER_EXECUTION
)

const (
	ScreenWidth    = 80
	DoubleLineChar = "="
	SingleLineChar = "-"

	ResourceFileExtension = ".resource"

	FlagLogLevel                   = "logLevel"
	FlagInputDir                   = "inputDir"
	FlagOutputDir                  = "outputDir"
	FlagExcludeFiles               = "excludeFiles"
	FlagPackageIdentifier          = "packageIdentifier"
	FlagResourceFileDelimiter      = "resourceFileDelimiter"
	FlagStringDelimiter            = "stringDelimiter"
	FlagSingleCommentDelimiter     = "singleCommentDelimiter"
	FlagBlockCommentDelimiterBegin = "blockCommentDelimiterBegin"
	FlagBlockCommentDelimiterEnd   = "blockCommentDelimiterEnd"

	UiLabelFilesToProcess                = "Files to process"
	UiTemplateProcessingFile             = "Processing file: %s"
	UiTemplateOutputFile                 = "Output file: %s"
	UiTemplateInputDir                   = "Input directory               : %s"
	UiTemplateOutputDir                  = "Output directory              : %s"
	UiTemplateExcludesFile               = "Exclude files                 : %s"
	UiTemplateLogLevel                   = "Log level                     : %s"
	UiTemplatePackageidentifier          = "Package identifier            : %s"
	UiTemplateResourceFileDelimiter      = "Resource file delimiter       : %s"
	UiTemplateStringDelimiter            = "String delimiter              : %+v"
	UiTemplateSingleLineDelimiter        = "Single line comment delimiter : %s"
	UiTemplateBlockCommentBeginDelimiter = "Block comment begin delimiter : %s"
	UiTemplateBlockCommentEndDelimiter   = "Block comment end delimiter   : %s"

	UiTemplateDirCollision = "Directory collision: %s overlaps with %s"

	ErrorTemplateInvocation            = "invocation error: [%v]"
	ErrorTemplateUndeterminedExecution = "undetermined execution error: [%v]"
	ErrorTemplateIo                    = "I/O error: [%v]"
	ErrorTemplateFileRead              = "error reading file: [%v]"
	ErrorTemplateFileWrite             = "error writing file: [%v]"
	ErrorTemplateParseError            = "parse error: [%+v]"
	ErrorTemplateTraverserExecution    = "error executing traverser: [%v]"

	LogTemplateCheckDirectoryCollision = "checking for directory collision: [%s: %s vs. %s: %s]"
	LogPrimaryDir                      = "primary directory"
	LogSecondaryDir                    = "secondary directory"
	LogTemplatePathsMatch              = "paths match: [%s]"
	LogTemplatePathCollision           = "path: [%s] collides with: [%s]"
	LogNoPathCollision                 = "paths do not collide"
	LogTemplateFileReadLine            = "read file line: [%s]"
	LogTemplateFileTrimmedLine         = "trimmed file line: [%s]"
	LogTemplateDelimiterNotDetected    = "delimiter not detected : [%s]"
	LogTemplateDelimiterDetected       = "delimiter detected : [%s]"
	LogTemplateDelimiterPosition       = "delimiter detected at position: [%s -> %d]"
	LogTemplateLiteralsDetected        = "literals detected: [%+v]"
	LogTemplateProcessingLiteral       = "processing literal: [%s]"
	LogTemplatePackage                 = "identified package: [%s]"
	LogTemplateFileOpen                = "opening file: [%s]"
	LogTemplateFileClose               = "closing file: [%s]"
	LogTemplateFileOutput              = "outputting file: [%s]"
	LogTemplateFileSkip                = "skipping file: [%s]"
	LogTemplateDirectorySkip           = "skipping directory: [%s]"
	LogTemplateConfig                  = "config: [%+v]"
	LogTemplateSettingLogLevel         = "setting log level: [%v]"
	LogTemplateSetLogLevel             = "set log level: [%v]"
	LogTemplateDirectoryExist          = "path: [%s]  is directory: [%t]"
)
