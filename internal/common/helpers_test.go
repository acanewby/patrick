package common

import (
	"testing"
)

type consolidateWhitespaceTest struct {
	name  string
	input string
	want  string
}

func TestConsolidateWhitespaceTableDriven(t *testing.T) {
	// Defining the tests
	const want = "Hello World"
	var tests = []consolidateWhitespaceTest{
		// the table itself
		{"spaces at beginning", "    Hello World", want},
		{"spaces at end", "Hello World   ", want},
		{"spaces in middle", "Hello    World", want},
		{"spaces everywhere", "   Hello    World   ", want},
		{"tabs in middle", "Hello \t \t World", want},
		{"newlines in middle", "Hello \n \n World", want},
		{"vertical tabs in middle", "Hello \v \v World", want},
		{"form feeds in middle", "Hello \f \f World", want},
		{"carriage returns in middle", "Hello \r \r World", want},
		{"NELs in middle", "Hello \u0085 \u0085 World", want},
		{"NBSPs in middle", "Hello \u00A0 \u00A0 World", want},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ConsolidateWhitespace(tt.input)
			if ans != tt.want {
				t.Errorf("got [%s], want [%s]", ans, tt.want)
			}
		})
	}
}

type directoryCollisionTest struct {
	name      string
	primary   string
	secondary string
	collision bool
	errStr    string
}

func TestDirectoryCollisionTableDriven(t *testing.T) {

	// Defining the tests
	const notADir = "/Users/Anewby/Dropbox/scratch/patrick/not-a-directory"
	const goodInputDir = "/Users/Anewby/Dropbox/scratch/patrick/input"
	const goodOutputDir = "/Users/Anewby/Dropbox/scratch/patrick/output"
	const childOutputDir = "/Users/Anewby/Dropbox/scratch/patrick/input/outputChild"
	const nonExistentDir = "/Users/Anewby/Dropbox/scratch/patrick/nope"
	var tests = []directoryCollisionTest{
		// the table itself
		{"directory does not exist", goodInputDir, nonExistentDir, false, "stat " + nonExistentDir + ": no such file or directory"},
		{"path is not a directory", notADir, nonExistentDir, false, "path: [" + notADir + "]  is directory: [false]"},
		{"directories match", goodInputDir, goodInputDir, true, ""},
		{"directories do not overlap", goodInputDir, goodOutputDir, false, ""},
		{"directory is a child", goodInputDir, childOutputDir, true, ""},
		{"directory is ancestor", childOutputDir, goodInputDir, true, ""},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collision, err := DirectoryCollision(tt.primary, tt.secondary)
			if collision != tt.collision {
				t.Errorf("got collision [%t], want [%t]", collision, tt.collision)
			}
			// We got an error when we didn't want one
			if err != nil && tt.errStr == "" {
				t.Errorf("got error [%+v], want [%s]", err, tt.errStr)
			}
			// We didn't get an error when we were expecting one
			if err == nil && tt.errStr != "" {
				t.Errorf("got error [%+v], want [%s]", err, tt.errStr)
			}
			// We got the wrong error
			if (err != nil && tt.errStr != "") && (err.Error() != tt.errStr) {
				t.Errorf("got error [%+v], want [%s]", err, tt.errStr)
			}
		})
	}
}
