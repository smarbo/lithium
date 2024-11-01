package tokens

type TokenType int

const (
	Exit TokenType = iota
	IntLit
	Semi
	OpenParen
	CloseParen
	Ident
	Let
	Eq
	Plus
)

type Token struct {
	Type  TokenType
	Value string
}
