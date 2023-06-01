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
	outputDir string
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
		patrick.Convert(constructConfig())
	},
}

func init() {

	var err error

	convertCmd.Flags().StringVarP(&outputDir, common.FlagOutputDir, "", "", "Output directory")

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
