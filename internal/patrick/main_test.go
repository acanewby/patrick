package patrick

import (
	"github.com/acanewby/patrick/internal/common"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		os.Exit(1)
	}

	exitCode := m.Run()

	if err := tearDown(); err != nil {
		os.Exit(1)
	}

	os.Exit(exitCode)

}

func setup() error {

	common.SetConfig(common.Config{
		InputDir:              "/Users/Anewby/Dropbox/scratch/patrick/input",
		OutputDir:             "/Users/Anewby/Dropbox/scratch/patrick/output",
		ExcludeFile:           "/Users/Anewby/Dropbox/scratch/patrick/exclude.list",
		LogLevel:              "debug",
		ResourceFileDelimiter: "|",
		LanguageConfig:        common.GoLanguageConfig(),
	})

	common.LogInfof("running tests: %s", time.Now())

	return nil
}

func tearDown() error {
	// tear down stuff here...
	return nil
}
