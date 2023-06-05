/*
Copyright © 2023 Adrian Newby <acanewby@yahoo.com>
*/
package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"github.com/acanewby/patrick/internal/patrick"
	"github.com/spf13/cobra"
	"os"
)

var (
	outputDir                  string
	packageIdentifier          string
	resourceFileDelimiter      string
	stringDelimiters           string
	singleCommentDelimiter     string
	blockCommentBeginDelimiter string
	blockCommentEndDelimiter   string
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Process files and apply string literal tokenizations",
	Long: `Processes the identified list of files and:

- identifies and tokenizes string literals
- outputs converted source files
- outputs associated resource files`,
	Run: func(cmd *cobra.Command, args []string) {
		common.SetConfig(constructConfig())
		patrick.Convert()
	},
}

func init() {

	var err error

	convertCmd.Flags().StringVarP(&outputDir, common.FlagOutputDir, "", "", "Output directory")
	convertCmd.Flags().StringVarP(&packageIdentifier, common.FlagPackageIdentifier, "", "package", "Keyword identifying package for this language system (default \"package\")")
	convertCmd.Flags().StringVarP(&resourceFileDelimiter, common.FlagResourceFileDelimiter, "", "|", "Delimiter separating resource id and value in resource file (default \",\")")
	convertCmd.Flags().StringVarP(&stringDelimiters, common.FlagStringDelimiters, "", "\",'", "Comma-separated list of string delimiters used by this language system (default\"\\\",'\")")
	convertCmd.Flags().StringVarP(&singleCommentDelimiter, common.FlagSingleCommentDelimiter, "", "//", "Single-line comment delimiter used by this language system (default\"//\")")
	convertCmd.Flags().StringVarP(&blockCommentBeginDelimiter, common.FlagBlockCommentDelimiterBegin, "", "/*", "Block comment begin delimiter used by this language system (default\"/*\")")
	convertCmd.Flags().StringVarP(&blockCommentEndDelimiter, common.FlagBlockCommentDelimiterEnd, "", "*/", "Block comment end delimiter used by this language system (default\"*/\")")

	if err = convertCmd.MarkFlagDirname(common.FlagOutputDir); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}
	if err = convertCmd.MarkFlagRequired(common.FlagOutputDir); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}

	rootCmd.AddCommand(convertCmd)

}
