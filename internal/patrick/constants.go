package patrick

import "github.com/acanewby/patrick/internal/common"

const (
	undefined codeState = iota + 1
	inNormalCode
	inCommentBlock
	inImportBlock
	onImportLine
	inConstBlock
	onConstLine
	onEmptyLine
)

const (
	InCodeStateNormal       = "In normal code"
	InCodeStateCommentBlock = "In comment block"
	InCodeStateImportBlock  = "In import block"
	InCodeStateConstBlock   = "In const block"
	OnCodeStateImportLine   = "On import line"
	OnCodeStateConstLine    = "On const line"
	OnCodeStateEmptyLine    = "On empty line"
)

func (typ codeState) String() string {

	var translation string

	switch typ {
	case inNormalCode:
		translation = InCodeStateNormal
	case inCommentBlock:
		translation = InCodeStateCommentBlock
	case inImportBlock:
		translation = InCodeStateImportBlock
	case inConstBlock:
		translation = InCodeStateConstBlock
	case onImportLine:
		translation = OnCodeStateImportLine
	case onConstLine:
		translation = OnCodeStateConstLine
	case onEmptyLine:
		translation = OnCodeStateEmptyLine
	default:
		translation = common.EnumUndefined
	}
	return translation
}

const (
	importKeyword       = "import"
	importBlockBegin    = importKeyword + " ("
	constKeyword        = "const"
	constBlockBegin     = constKeyword + " ("
	importConstBlockEnd = ")"
)
