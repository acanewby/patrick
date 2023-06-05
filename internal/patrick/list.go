package patrick

import (
	"fmt"
	"github.com/acanewby/patrick/internal/common"
	"os"
)

func List() {

	cfg := common.GetConfig()

	excludeList := setupRun(cfg)
	var err error

	// Process the targeted files as a slice
	if err = common.TraverseFilteredDirectoryTree(cfg.InputDir, excludeList, listFilename); err != nil {
		common.LogErrorf(common.ErrorTemplateTraverserExecution, err)
		os.Exit(common.EXIT_CODER_TRAVERSER_EXECUTION)

	}

}

// listFilename implements traverseWorker
// It writes the value of path to the console
func listFilename(path string) error {
	fmt.Println(fmt.Sprintf(common.UiTemplateProcessingFile, path))
	return nil
}
