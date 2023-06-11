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
	outputDir             string
	resourceFileDelimiter string
	overwriteOutput       bool
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Process files and apply string literal tokenizations",
	Long: `
Processes the identified list of files and:

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

	convertCmd.Flags().StringVarP(&outputDir, common.FlagOutputDir, "", "", "output directory")
	convertCmd.Flags().StringVarP(&resourceFileDelimiter, common.FlagResourceFileDelimiter, "", "|", "delimiter separating resource id and value in resource file")
	convertCmd.Flags().BoolVarP(&overwriteOutput, common.FlagOverwriteOutput, "", false, "replace contents of outputDir")

	// Make these flags if we ever make a non Go-specific version of patrick
	//convertCmd.Flags().StringVarP(&packageIdentifier, common.FlagPackageIdentifier, "", "package", "keyword identifying package for this language system")
	//convertCmd.Flags().StringVarP(&stringDelimiter, common.FlagStringDelimiter, "", "\"", "string delimiter used by this language system")
	//convertCmd.Flags().StringVarP(&singleCommentDelimiter, common.FlagSingleCommentDelimiter, "", "//", "single-line comment delimiter used by this language system")
	//convertCmd.Flags().StringVarP(&blockCommentBeginDelimiter, common.FlagBlockCommentDelimiterBegin, "", "/*", "block comment begin delimiter used by this language system")
	//convertCmd.Flags().StringVarP(&blockCommentEndDelimiter, common.FlagBlockCommentDelimiterEnd, "", "*/", "block comment end delimiter used by this language system")

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
