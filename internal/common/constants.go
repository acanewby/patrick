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

	ResourceFileExtension                     = ".resource"
	ResourceFunctionTemplateSubstitutionToken = "%%RESOURCE_INDEX%%"

	EnumUndefined = "ENUM_UNDEFINED"

	FlagLogLevel                 = "logLevel"
	FlagInputDir                 = "inputDir"
	FlagOutputDir                = "outputDir"
	FlagOverwriteOutput          = "overwriteOutput"
	FlagExcludeFiles             = "excludeFiles"
	FlagResourceFileDelimiter    = "resourceFileDelimiter"
	FlagResourceIndexStart       = "resourceIndexStart"
	FlagResourceIndexZeroPad     = "resourceIndexZeroPad"
	FlagResourceTokenPrefix      = "resourceTokenPrefix"
	FlagResourceFunctionTemplate = "resourceFunctionTemplate"

	UiLabelFilesToProcess                = "Files to process"
	UiTemplateProcessingFile             = "Processing file: %s"
	UiTemplateOutputFile                 = "Output file: %s"
	UiTemplateInputDir                   = "Input directory               : %s"
	UiTemplateOutputDir                  = "Output directory              : %s"
	UiTemplateOverwriteOutput            = "Overwrite output              : %t"
	UiTemplateExcludesFile               = "Exclude files                 : %s"
	UiTemplateLogLevel                   = "Log level                     : %s"
	UiTemplatePackageidentifier          = "Package identifier            : %s"
	UiTemplateResourceFileDelimiter      = "Resource file delimiter       : %s"
	UiTemplateResourceIndexStart         = "Resource index start          : %d"
	UiTemplateResourceIndexZeroPad       = "Resource index zero pad       : %d"
	UiTemplateResourceTokenPrefix        = "Resource token prefix         : %s"
	UiTemplateResourceFunctionTemplate   = "Resource function template    : %s"
	UiTemplateStringDelimiter            = "String delimiter              : %+v"
	UiTemplateSingleLineDelimiter        = "Single line comment delimiter : %s"
	UiTemplateBlockCommentBeginDelimiter = "Block comment begin delimiter : %s"
	UiTemplateBlockCommentEndDelimiter   = "Block comment end delimiter   : %s"
	UiTemplateImportKeyword              = "Import keyword                : %s"
	UiTemplateImportBlockDelimiters      = "Import block delimiters       : %s ... %s"
	UiTemplateConstKeyword               = "Const keyword                 : %s"
	UiTemplateConstBlockDelimiters       = "Const block delimiters        : %s ... %s"

	UiTemplateDirCollision = "Directory collision: %s overlaps with %s"

	ErrorTemplateResourceFunctionMissingSubToken = "resource function template does not include %s : [%s]"
	ErrorTemplateInvocation                      = "invocation error: [%v]"
	ErrorTemplateUndeterminedExecution           = "undetermined execution error: [%v]"
	ErrorTemplateIo                              = "I/O error: [%v]"
	ErrorTemplateOutputDirAlreadyExists          = "output directory already exists: [%s]"
	ErrorTemplateFileRead                        = "error reading file: [%v]"
	ErrorTemplateFileWrite                       = "error writing file: [%v]"
	ErrorTemplateParseError                      = "parse error: [%+v]"
	ErrorTemplateTraverserExecution              = "error executing traverser: [%v]"

	LogTemplateCheckDirectoryCollision = "checking for directory collision: [%s: %s vs. %s: %s]"
	LogTemplateExclusions              = "exclude files: [%+v]"
	LogTemplateTextFileContents        = "text file contents: [%+v]"
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
	LogLiteralsNotDetected             = "no literals detected"
	LogTemplateProcessingLiteral       = "processing literal: [%s]"
	LogTemplatePackage                 = "identified package: [%s]"
	LogTemplateDirCreate               = "creating directory: [%s]"
	LogTemplateDirRemove               = "deleting directory: [%s]"
	LogTemplateFileOpen                = "opening file: [%s]"
	LogTemplateFileClose               = "closing file: [%s]"
	LogTemplateFileOutput              = "outputting file: [%s]"
	LogTemplateFileSkip                = "skipping file: [%s]"
	LogTemplateDirectorySkip           = "skipping directory: [%s]"
	LogTemplateConfig                  = "config: [%+v]"
	LogTemplateSettingLogLevel         = "setting log level: [%v]"
	LogTemplateSetLogLevel             = "set log level: [%v]"
	LogTemplateFileNotDirectory        = "path: [%s]  is file, not directory"
	LogTemplateDirectoryNotFile        = "path: [%s]  is directory, not file"
	LogTemplateDirectoryExist          = "path: [%s]  is directory: [%t]"
	LogTemplateFileExist               = "path: [%s]  is file: [%t]"
	LogTemplateShouldParse             = "code state: [%s -> %s], should parse: [%t]"
)
