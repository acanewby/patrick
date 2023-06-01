/*
Copyright © 2023 Adrian Newby <acanewby@yahoo.com>
*/
package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"os"

	"github.com/spf13/cobra"
)

var (
	version           = "x.x.x"
	excludedNamesFile string
	inputDir          string
	logLevel          string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "patrick",
	Version: version,
	Short:   "A utility to extract string literals from source code",
	Long: `patrick is a software engineering utility that assists with application globalization/localization.

It parses specified collections of source code files and identifies string literals, which it then tokenizes,
producing tokenized source files and associated resource data files, suitable for most resource management approaches.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateUndeterminedExecution, err))
		os.Exit(common.EXIT_CODE_UNDETERMINED_ERROR)
	}
}

func init() {

	var err error

	rootCmd.PersistentFlags().StringVarP(&logLevel, common.FlagLogLevel, "", "", "log level (debug,info,warn,error,fatal)")

	rootCmd.PersistentFlags().StringVarP(&excludedNamesFile, common.FlagExcludeFiles, "", "", "file containing base filenames to exclude - one per line")

	if err = rootCmd.MarkPersistentFlagFilename(common.FlagExcludeFiles); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}

	rootCmd.PersistentFlags().StringVarP(&inputDir, common.FlagInputDir, "", "", "input directory")

	if err = rootCmd.MarkPersistentFlagDirname(common.FlagInputDir); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}
	if err = rootCmd.MarkPersistentFlagRequired(common.FlagInputDir); err != nil {
		fmt.Println(fmt.Sprintf(common.ErrorTemplateInvocation, err))
		os.Exit(common.EXIT_CODE_INVOCATION_ERROR)
	}

}
