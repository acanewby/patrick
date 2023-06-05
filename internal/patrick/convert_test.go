package patrick

import (
	"testing"
)

type semanticLineTest struct {
	name  string
	input string
	want  string
}

func TestSemanticLineTableDriven(t *testing.T) {
	// Defining the tests
	const want = "package test"
	var tests = []semanticLineTest{
		// the table itself
		{"single line comment at end", "package test // single line", want},
		{"block comment begin at end", "package test /* block", want},
		{"block comment embedded", "package /* cheeky little mid-line block */ test", want},
		{"back-to-back block comments", "   */  package test /* block", want},
		{"block end comment", "   */  package test   ", want},
		{"block end comment and single-line", "   */  package test // single", want},
		{"block comment embedded with extra whitespace", "package /* cheeky little mid-line block */    test", want},
		{"single line comment embedded in block comment", "package test /* how about a single line comment //  in the middle of a block", want},
		{"block comment embedded in single line comment", "package test // and a /* block */ too", want},
		{"block comment embedded with single-line comment at end", "package /* can we handle ... */ test // this ... ?", want},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := semanticLine(tt.input)
			if ans != tt.want {
				t.Errorf("got [%s], want [%s]", ans, tt.want)
			}
		})
	}
}
