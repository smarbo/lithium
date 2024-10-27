package unicodeext

import "unicode"

var (
	IsDigit  = unicode.IsDigit
	IsLetter = unicode.IsLetter
	IsSpace  = unicode.IsSpace
	IsMark   = unicode.IsMark
	IsUpper  = unicode.IsUpper
	IsLower  = unicode.IsLower
	IsOneOf  = unicode.IsOneOf
	IsPrint  = unicode.IsPrint
	IsPunct  = unicode.IsPunct
	IsTitle  = unicode.IsTitle
	IsNumber = unicode.IsNumber
	IsSymbol = unicode.IsSymbol
)

// Returns true if the rune is either a letter or digit
func IsAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
