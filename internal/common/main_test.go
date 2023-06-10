package common

import (
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

	SetConfig(Config{
		InputDir:                   "/Users/Anewby/Dropbox/scratch/patrick/input",
		OutputDir:                  "/Users/Anewby/Dropbox/patrick/scratch/output",
		ExcludeFile:                "/Users/Anewby/Dropbox/patrick/scratch/exclude.list",
		LogLevel:                   "debug",
		PackageIdentifier:          "package",
		ResourceFileDelimiter:      "|",
		StringDelimiter:            "\"",
		SingleLineCommentDelimiter: "//",
		BlockCommentBeginDelimiter: "/*",
		BlockCommentEndDelimiter:   "*/",
	})

	LogInfof("running tests: %s", time.Now())

	return nil
}

func tearDown() error {
	// tear down stuff here...
	return nil
}
