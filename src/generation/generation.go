package generation

import (
	"log"
	"strconv"

	p "github.com/smarbo/lithium/src/parser"
)

func (g *Generator) GenExpr(expr p.NodeExpr) {
	switch v := expr.Var.(type) {
	case p.NodeExprIntLit:
		g.output += "    mov rax, " + v.IntLit.Value + "\n"
		g.push("rax")
	case p.NodeExprIdent:
		if val, exists := g.vars[v.Ident.Value]; exists {
			// identifier does exist
			var offset string
			offset += "QWORD [rsp + " + strconv.Itoa((g.stack_size-val.stack_loc-1)*8) + "]\n"
			g.push(offset)
		} else {
			log.Fatal("lithium compilation error: undeclared identifier: " + v.Ident.Value)
		}
	case p.NodeBinExpr:
		log.Fatal("not implemented yet (binary expresion)")
	}
}

func (g *Generator) GenStmt(stmt p.NodeStmt) {
	switch v := stmt.Var.(type) {
	case p.NodeStmtLet:
		if _, exists := g.vars[v.Ident.Value]; exists {
			log.Fatal("lithium compilation error: identifier already used: " + v.Ident.Value)
		}

		g.vars[v.Ident.Value] = Var{stack_loc: g.stack_size}
		g.GenExpr(v.Expr)
		break
	case p.NodeStmtExit:
		g.GenExpr(v.Expr)
		g.output += "    mov rax, 60\n"
		g.pop("rdi")
		g.output += "    syscall\n"
	}
}

func (g *Generator) GenProg() string {
	g.output += "; AUTOGENERATED BY LITHIUM COMPILER ;\nglobal _start\n_start:\n"

	for _, stmt := range g.prog.Stmts {
		g.GenStmt(stmt)
	}

	g.output += "    ; AUTOGENERATED EXIT CODE ;\n"
	g.output += "    mov rax, 60\n"
	g.output += "    mov rdi, 0\n"
	g.output += "    syscall\n"
	return g.output
}

func (g *Generator) push(reg string) {
	g.output += "    push " + reg + "\n"
	g.stack_size++
}

func (g *Generator) pop(reg string) {
	g.output += "    pop " + reg + "\n"
	g.stack_size--
}

func Init(prog p.NodeProg) *Generator {
	return &Generator{
		prog:       prog,
		output:     "",
		stack_size: 0,
		vars:       make(map[string]Var),
	}
}

type Var struct {
	stack_loc int
}

type Generator struct {
	prog       p.NodeProg
	output     string
	stack_size int
	vars       map[string]Var
}
