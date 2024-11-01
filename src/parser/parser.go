package parser

import (
	"errors"
	"log"

	t "github.com/smarbo/lithium/src/tokens"
)

type Parser struct {
	tokens []t.Token
	index  int
}

func (p *Parser) peek(offset int) *t.Token {
	if p.index+offset >= len(p.tokens) {
		return nil
	} else {
		tok := p.tokens[p.index+offset]
		return &tok
	}
}

func (p *Parser) consume() t.Token {
	p.index++
	return p.tokens[p.index-1]
}

func (p *Parser) ParseBinExpr() *NodeBinExpr {
	if lhs := p.ParseExpr(); lhs != nil {
		bin_expr := NodeBinExpr{}
		if p.peek(0) != nil && p.peek(0).Type == t.Plus {
			bin_expr_add := NodeBinExprAdd{lhs: *lhs}
			p.consume()
			if rhs := p.ParseExpr(); rhs != nil {
				bin_expr_add.rhs = *rhs
				bin_expr.Var = bin_expr_add
				return &bin_expr
			} else {
				log.Fatal("lithium compilation error: expected expression")
				return nil
			}
		} else {
			log.Fatal("lithium compilation error: unsupported binary operator")
			return nil
		}
	} else {
		return nil
	}
}

func (p *Parser) ParseTerm() *NodeTerm {
	if p.peek(0) != nil && p.peek(0).Type == t.IntLit {
		return &NodeTerm{Var: NodeTermIntLit{IntLit: p.consume()}}
	} else if p.peek(0) != nil && p.peek(0).Type == t.Ident {
		return &NodeTerm{Var: NodeTermIdent{Ident: p.consume()}}
	} else {
		return nil
	}
}

func (p *Parser) ParseExpr() *NodeExpr {
	if term := p.ParseTerm(); term != nil {
		return nil
	} else {
		return nil
	}

	//else if bin_expr := p.ParseBinExpr(); bin_expr != nil {
	//return &NodeExpr{Var: bin_expr}
	//} else {
	//return nil
	//}
}

func (p *Parser) ParseStmt() (error, *NodeStmt) {
	if p.peek(0).Type == t.Exit && p.peek(1) != nil && p.peek(1).Type == t.OpenParen {
		p.consume()
		p.consume()
		var stmt_exit NodeStmtExit
		if node_expr := p.ParseExpr(); node_expr != nil {
			stmt_exit = NodeStmtExit{Expr: *node_expr}
		} else {
			return errors.New("invalid expression"), nil
		}

		if p.peek(0) != nil && p.peek(0).Type == t.CloseParen {
			p.consume()
		} else {
			return errors.New("expected `)`"), nil
		}
		if p.peek(0) != nil && p.peek(0).Type == t.Semi {
			p.consume()
		} else {
			return errors.New("expected `;`"), nil
		}
		return nil, &NodeStmt{Var: stmt_exit}
	} else if p.peek(0) != nil && p.peek(0).Type == t.Let &&
		p.peek(1) != nil && p.peek(1).Type == t.Ident &&
		p.peek(2) != nil && p.peek(2).Type == t.Eq {
		p.consume()
		stmt_let := NodeStmtLet{Ident: p.consume()}
		p.consume()
		if expr := p.ParseExpr(); expr != nil {
			stmt_let.Expr = *expr
		} else {
			return errors.New("invalid expression"), nil
		}
		if p.peek(0) != nil && p.peek(0).Type == t.Semi {
			p.consume()
		} else {
			return errors.New("expected `;`"), nil
		}

		return nil, &NodeStmt{Var: stmt_let}
	} else {
		return errors.New("unexpected statement type"), nil
	}
}

func (p *Parser) ParseProg() (error, *NodeProg) {
	var prog NodeProg
	for p.peek(0) != nil {
		if _, stmt := p.ParseStmt(); stmt != nil {
			prog.Stmts = append(prog.Stmts, *stmt)
		} else {
			return errors.New("invalid statement"), nil
		}
	}

	return nil, &prog
}

func Init(tokens []t.Token) *Parser {
	return &Parser{
		tokens: tokens,
		index:  0,
	}
}
