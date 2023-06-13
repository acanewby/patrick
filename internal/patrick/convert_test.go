package patrick

import (
	"testing"
)

type semanticLineTest struct {
	name                          string
	input                         string
	semanticResultExpected        string
	blockCommentBeganFlagExpected bool
	blockCommentEndedFlagExpected bool
}

func TestSemanticLineTableDriven(t *testing.T) {
	// Defining the tests
	const semanticResult = "package test"
	var tests = []semanticLineTest{
		// the table itself
		{"single line comment at end", "package test // single line",
			semanticResult, false, false},
		{"block comment begin at end", "package test /* block",
			semanticResult, true, false},
		{"block comment embedded", "package /* cheeky little mid-line block */ test",
			semanticResult, false, false},
		{"back-to-back block comments", "   */  package test /* block",
			semanticResult, false, false},
		{"block end comment", "   */  package test   ",
			semanticResult, false, true},
		{"block end comment and single-line", "   */  package test // single",
			semanticResult, false, true},
		{"block comment embedded with extra whitespace", "package /* cheeky little mid-line block */    test",
			semanticResult, false, false},
		{"single line comment embedded in block comment", "package test /* how about a single line comment //  in the middle of a block",
			semanticResult, true, false},
		{"block comment embedded in single line comment", "package test // and a /* block */ too",
			semanticResult, false, false},
		{"block comment embedded with single-line comment at end", "package /* can we handle ... */ test // this ... ?",
			semanticResult, false, false},
		{"block comment embedded with block comment begin at end", "package /* can we handle ... */ test /* this ... ?",
			semanticResult, true, false},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			semanticResult, blockCommentBegan, blockCommentEnded := semanticLine(tt.input)
			if semanticResult != tt.semanticResultExpected {
				t.Errorf("got [%s], want [%s]", semanticResult, tt.semanticResultExpected)
			}
			if blockCommentEnded != tt.blockCommentEndedFlagExpected {
				t.Errorf("got [%t], want [%t]", blockCommentEnded, tt.blockCommentEndedFlagExpected)
			}
			if blockCommentBegan != tt.blockCommentBeganFlagExpected {
				t.Errorf("got [%t], want [%t]", blockCommentBegan, tt.blockCommentBeganFlagExpected)
			}
		})
	}
}
