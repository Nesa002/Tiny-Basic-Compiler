package optimizer

import (
	"tiny-basic/src/ast"
)

func Optimize(program *ast.Program) *ast.Program {
	newStmts := []ast.Statement{}

	for _, stmt := range program.Statements {
		newStmts = append(newStmts, stmt)

		if _, ok := stmt.(*ast.EndStatement); ok {
			break
		}
	}

	return &ast.Program{
		Statements: newStmts,
	}
}
