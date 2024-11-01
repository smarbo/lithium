package parser

import t "github.com/smarbo/lithium/src/tokens"

type NodeExprI interface {
	IsNodeExpr()
}

type NodeTermI interface {
	IsNodeTerm()
}

type NodeTermIntLit struct {
	IntLit t.Token
}

type NodeTermIdent struct {
	Ident t.Token
}

func (n NodeTerm) IsNodeExpr()    {}
func (n NodeBinExpr) IsNodeExpr() {}

func (n NodeTermIntLit) IsNodeTerm() {}
func (n NodeTermIdent) IsNodeTerm()  {}

type NodeExpr struct {
	Var NodeExprI
}

type NodeTerm struct {
	Var NodeTermI
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

type NodeBinExprAdd struct {
	lhs NodeExpr
	rhs NodeExpr
}

type NodeBinExprMulti struct {
	lhs NodeExpr
	rhs NodeExpr
}

type BinExprI interface {
	IsBinExpr()
}

func (b NodeBinExprAdd) IsBinExpr()   {}
func (b NodeBinExprMulti) IsBinExpr() {}

type NodeBinExpr struct {
	Var BinExprI
}
