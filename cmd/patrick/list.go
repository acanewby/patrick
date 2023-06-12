/*
Copyright © 2023 Adrian Newby <acanewby@yahoo.com>
*/
package patrick

import (
	"github.com/acanewby/patrick/internal/common"
	"github.com/acanewby/patrick/internal/patrick"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the files that will be targeted for processing",
	Long: `
Lists all files that will be targeted for processing, taking into account inclusion and exclusion criteria.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		common.SetConfig(constructConfig())
	},
	Run: func(cmd *cobra.Command, args []string) {
		patrick.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
