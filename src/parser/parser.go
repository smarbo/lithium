package parser

import (
	"errors"

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

func (p *Parser) ParseExpr() *NodeExpr {
	if p.peek(0) != nil && p.peek(0).Type == t.IntLit {
		return &(NodeExpr{Var: NodeExprIntLit{IntLit: p.consume()}})
	} else if p.peek(0) != nil && p.peek(0).Type == t.Ident {
		return &(NodeExpr{Var: NodeExprIdent{Ident: p.consume()}})
	} else {
		return nil
	}
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
