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
	outputDir       string
	overwriteOutput bool

	resourceFileDelimiter    string
	resourceIndexStart       uint64
	resourceIndexZeroPad     uint8
	resourceTokenPrefix      string
	resourceFunctionTemplate string
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
	PreRun: func(cmd *cobra.Command, args []string) {
		common.SetConfig(constructConfig())
	},
	Run: func(cmd *cobra.Command, args []string) {
		patrick.Convert()
	},
}

func init() {

	var err error

	convertCmd.Flags().StringVarP(&outputDir, common.FlagOutputDir, "", "", "output directory")
	convertCmd.Flags().BoolVarP(&overwriteOutput, common.FlagOverwriteOutput, "", false, "replace contents of outputDir")

	convertCmd.Flags().StringVarP(&resourceFileDelimiter, common.FlagResourceFileDelimiter, "", " = ", "delimiter separating resource id and value in resource file")

	convertCmd.Flags().Uint64VarP(&resourceIndexStart, common.FlagResourceIndexStart, "", 10000, "starting value for sequentially-numbered resource tokens")
	convertCmd.Flags().Uint8VarP(&resourceIndexZeroPad, common.FlagResourceIndexZeroPad, "", 8, "width of zero-padded resource token index number")
	convertCmd.Flags().StringVarP(&resourceTokenPrefix, common.FlagResourceTokenPrefix, "", "Resource_", "prefix string to be prepended to resource tokens")
	convertCmd.Flags().StringVarP(&resourceFunctionTemplate, common.FlagResourceFunctionTemplate, "", "", fmt.Sprintf("resource function to be substituted into source code in place of each identified string literal. Must include %s, which will be the placeholder for the generated resource token", common.ResourceFunctionTemplateSubstitutionToken))

	if err = convertCmd.MarkFlagDirname(common.FlagOutputDir); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}
	if err = convertCmd.MarkFlagRequired(common.FlagOutputDir); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}
	if err = convertCmd.MarkFlagRequired(common.FlagResourceFunctionTemplate); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}

	rootCmd.AddCommand(convertCmd)

}
