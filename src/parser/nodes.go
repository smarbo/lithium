package parser

import t "github.com/smarbo/lithium/src/tokens"

type NodeExprI interface {
	IsNodeExpr()
}

type NodeExprIntLit struct {
	IntLit t.Token
}

type NodeExprIdent struct {
	Ident t.Token
}

func (n NodeExprIntLit) IsNodeExpr() {}
func (n NodeExprIdent) IsNodeExpr()  {}

type NodeExpr struct {
	Var NodeExprI
}

type NodeStmtI interface {
	IsNodeStmt()
}

type NodeStmtExit struct {
	Expr NodeExpr
}

type NodeStmtLet struct {
	Ident t.Token
	Expr  NodeExpr
}

func (n NodeStmtExit) IsNodeStmt() {}
func (n NodeStmtLet) IsNodeStmt()  {}

type NodeStmt struct {
	Var NodeStmtI
}

type NodeProg struct {
	Stmts []NodeStmt
}
