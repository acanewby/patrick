package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"os"
	"regexp"
)

func setupRun(cfg common.Config) []string {

	var err error
	var excludeList []string

	if cfg.LogLevel != "" {
		common.SetLogLevel(cfg.LogLevel)
	}

	dumpConfig(cfg)

	// Get a slice of excludable filenames
	if cfg.ExcludeFile != "" {
		if excludeList, err = common.GetTextFileContents(cfg.ExcludeFile); err != nil {
			msg := fmt.Sprintf(common.ErrorTemplateFileRead, err)
			common.LogErrorf(msg)
			fmt.Println(msg)
			os.Exit(common.EXIT_CODE_IO_ERROR)
		}
	}

	common.LogInfof(common.LogTemplateExclusions, excludeList)
	return excludeList
}

func dumpConfig(cfg common.Config) {
	common.DoubleLineToConsole()

	common.LogDebugf(common.LogTemplateConfig, cfg)

	fmt.Println(fmt.Sprintf(common.UiTemplateInputDir, cfg.InputDir))
	fmt.Println(fmt.Sprintf(common.UiTemplateOutputDir, cfg.OutputDir))
	fmt.Println(fmt.Sprintf(common.UiTemplateExcludesFile, cfg.ExcludeFile))
	fmt.Println(fmt.Sprintf(common.UiTemplateLogLevel, cfg.LogLevel))
	fmt.Println(fmt.Sprintf(common.UiTemplatePackageidentifier, cfg.LanguageConfig.PackageIdentifier))
	fmt.Println(fmt.Sprintf(common.UiTemplateResourceFileDelimiter, cfg.ResourceFileDelimiter))

	fmt.Println(fmt.Sprintf(common.UiTemplateStringDelimiter, cfg.LanguageConfig.StringDelimiter))

	fmt.Println(fmt.Sprintf(common.UiTemplateSingleLineDelimiter, cfg.LanguageConfig.SingleLineCommentDelimiter))
	fmt.Println(fmt.Sprintf(common.UiTemplateBlockCommentBeginDelimiter, cfg.LanguageConfig.BlockCommentBeginDelimiter))
	fmt.Println(fmt.Sprintf(common.UiTemplateBlockCommentEndDelimiter, cfg.LanguageConfig.BlockCommentEndDelimiter))

	fmt.Println(fmt.Sprintf(common.UiTemplateImportKeyword, cfg.LanguageConfig.ImportKeyword))
	fmt.Println(fmt.Sprintf(common.UiTemplateImportBlockDelimiters, cfg.LanguageConfig.ImportBlockBegin, cfg.LanguageConfig.ImportBlockEnd))
	fmt.Println(fmt.Sprintf(common.UiTemplateConstKeyword, cfg.LanguageConfig.ConstKeyword))
	fmt.Println(fmt.Sprintf(common.UiTemplateConstBlockDelimiters, cfg.LanguageConfig.ConstBlockBegin, cfg.LanguageConfig.ConstBlockEnd))

	common.DoubleLineToConsole()
}

func extractStringLiterals(line string) []string {

	cfg := common.GetConfig()

	// Find instances of non-string-delimiter characters between pairs of string delimiters
	r := regexp.MustCompile(fmt.Sprintf(`(\%s[^%s]*\%s)`, cfg.LanguageConfig.StringDelimiter, cfg.LanguageConfig.StringDelimiter, cfg.LanguageConfig.StringDelimiter))

	// Get the list of identified literals
	literals := r.FindAllString(line, -1)

	if literals == nil {
		common.LogInfof(common.LogLiteralsNotDetected)
	} else {
		common.LogInfof(common.LogTemplateLiteralsDetected, literals)
	}

	return literals
}
