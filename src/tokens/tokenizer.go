package tokens

import (
	"errors"

	uni "github.com/smarbo/lithium/src/unicode_ext"
)

type Tokenizer struct {
	src   string
	index int
}

func (t *Tokenizer) peek(offset int) *rune {
	if t.index+offset > len(t.src) {
		return nil
	} else {
		r := rune(t.src[t.index])
		return &r
	}
}

func (t *Tokenizer) consume() rune {
	t.index++
	return rune(t.src[t.index-1])
}

func (t *Tokenizer) Tokenize() (error, []Token) {
	var tokens []Token
	var buf string

	for t.peek(1) != nil {
		if uni.IsLetter(*t.peek(1)) {
			buf += string(t.consume())
			for t.peek(1) != nil && uni.IsAlnum(*t.peek(1)) {
				buf += string(t.consume())
			}

			if buf == "exit" {
				tokens = append(tokens, Token{
					Type: Exit,
				})
				buf = ""
				continue
			} else if buf == "let" {
				tokens = append(tokens, Token{
					Type: Let,
				})
				buf = ""
				continue
			} else {
				tokens = append(tokens, Token{
					Type:  Ident,
					Value: buf,
				})
				buf = ""
				continue
			}
		} else if uni.IsDigit(*t.peek(1)) {
			buf += string(t.consume())
			for t.peek(1) != nil && uni.IsDigit(*t.peek(1)) {
				buf += string(t.consume())
			}
			tokens = append(tokens, Token{
				Type:  IntLit,
				Value: buf,
			})
			buf = ""
			continue
		} else if *(t.peek(1)) == ';' {
			t.consume()
			tokens = append(tokens, Token{
				Type: Semi,
			})
			continue
		} else if *(t.peek(1)) == '=' {
			t.consume()
			tokens = append(tokens, Token{
				Type: Eq,
			})
			continue
		} else if *(t.peek(1)) == '(' {
			t.consume()
			tokens = append(tokens, Token{
				Type: OpenParen,
			})
			continue
		} else if *(t.peek(1)) == ')' {
			t.consume()
			tokens = append(tokens, Token{
				Type: CloseParen,
			})
			continue
		} else if uni.IsSpace(*t.peek(1)) {
			t.consume()
			continue
		} else {
			return errors.New("invalid token"), []Token{}
		}
	}

	t.index = 0
	return nil, tokens
}

func Init(src string) *Tokenizer {
	return &Tokenizer{
		src:   src,
		index: 0,
	}
}
