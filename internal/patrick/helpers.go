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

    // Get a slice of excludable filenames
    if cfg.ExcludeFile != "" {
        if excludeList, err = common.GetTextFileContents(cfg.ExcludeFile); err != nil {
            common.LogErrorf(common.ErrorTemplateFileRead, err)
            os.Exit(common.EXIT_CODE_IO_ERROR)
        }
    }

    dumpConfig(cfg)

    return excludeList
}

func dumpConfig(cfg common.Config) {
    common.DoubleLineToConsole()

    common.LogInfof(common.LogTemplateConfig, cfg)

    fmt.Println(fmt.Sprintf(common.UiTemplateInputDir, cfg.InputDir))
    fmt.Println(fmt.Sprintf(common.UiTemplateOutputDir, cfg.OutputDir))
    fmt.Println(fmt.Sprintf(common.UiTemplateExcludesFile, cfg.ExcludeFile))
    fmt.Println(fmt.Sprintf(common.UiTemplateLogLevel, cfg.LogLevel))
    fmt.Println(fmt.Sprintf(common.UiTemplatePackageidentifier, cfg.PackageIdentifier))
    fmt.Println(fmt.Sprintf(common.UiTemplateResourceFileDelimiter, cfg.ResourceFileDelimiter))
    fmt.Println(fmt.Sprintf(common.UiTemplateStringDelimiter, cfg.StringDelimiter))
    fmt.Println(fmt.Sprintf(common.UiTemplateSingleLineDelimiter, cfg.SingleLineCommentDelimiter))
    fmt.Println(fmt.Sprintf(common.UiTemplateBlockCommentBeginDelimiter, cfg.BlockCommentBeginDelimiter))
    fmt.Println(fmt.Sprintf(common.UiTemplateBlockCommentEndDelimiter, cfg.BlockCommentEndDelimiter))

    common.DoubleLineToConsole()
}

func extractStringLiterals(line string) []string {

    cfg := common.GetConfig()

    // Find instances of non-string-delimiter characters between pairs of string delimiters
    r := regexp.MustCompile(fmt.Sprintf(`(\%s[^%s]*\%s)`, cfg.StringDelimiter, cfg.StringDelimiter, cfg.StringDelimiter))

    // Get the list of identified literals
    literals := r.FindAllString(line, -1)
    
    if literals == nil {
        common.LogDebugf(common.LogLiteralsNotDetected)
    } else {
        common.LogDebugf(common.LogTemplateLiteralsDetected, literals)
    }

    return literals
}
