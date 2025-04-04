package codegen

import (
	"fmt"
	"strings"
	"tiny-basic/src/ast"
)

type CodeGenerator struct {
	builder strings.Builder
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

func (cg *CodeGenerator) Generate(program *ast.Program) string {
	for _, stmt := range program.Statements {
		cg.builder.WriteString(cg.generateStatement(stmt) + "\n")
	}

	return cg.builder.String()
}

func (cg *CodeGenerator) generateStatement(stmt ast.Statement) string {
	switch stmt := stmt.(type) {
	case *ast.PrintStatement:
		return cg.generatePrintStatement(stmt)
	case *ast.IfStatement:
		return cg.generateIfStatement(stmt)
	case *ast.LetStatement:
		return cg.generateLetStatement(stmt)
	case *ast.AssignmentStatement:
		return cg.generateAssignmentStatement(stmt)
	case *ast.EndStatement:
		return "process.exit(0);"
	case *ast.CommentStatement:
		return "// " + stmt.Text
	default:
		return ""
	}
}

func (cg *CodeGenerator) generatePrintStatement(stmt *ast.PrintStatement) string {
	return "console.log(" + cg.generateExpression(stmt.Expression, false) + ");"
}

func (cg *CodeGenerator) generateIfStatement(stmt *ast.IfStatement) string {
	condition := cg.generateExpression(stmt.Condition, false)
	thenBranch := cg.generateStatement(stmt.ThenBranch)
	elseBranch := ""

	if stmt.ElseBranch != nil {
		elseBranch = fmt.Sprintf(" else {\n\t%s\n}", cg.generateStatement(stmt.ElseBranch))
	}

	return fmt.Sprintf("if (%s) {\n\t%s\n}%s", condition, thenBranch, elseBranch)
}

func (cg *CodeGenerator) generateLetStatement(stmt *ast.LetStatement) string {
	return fmt.Sprintf("let %s = %s;", stmt.Identifier.Name, cg.generateExpression(stmt.Value, false))
}

func (cg *CodeGenerator) generateAssignmentStatement(stmt *ast.AssignmentStatement) string {
	return fmt.Sprintf("%s = %s;", stmt.Identifier.Name, cg.generateExpression(stmt.Value, false))
}

func (cg *CodeGenerator) generateExpression(expr ast.Expression, addParentheses bool) string {
	switch expr := expr.(type) {
	case *ast.Identifier:
		return expr.Name
	case *ast.FloatLiteral:
		return fmt.Sprintf("%v", expr.Value)
	case *ast.IntegerLiteral:
		return fmt.Sprintf("%v", expr.Value)
	case *ast.BinaryExpression:
		left := cg.generateExpression(expr.Left, true)
		right := cg.generateExpression(expr.Right, true)
		if addParentheses {
			return fmt.Sprintf("(%s %s %s)", left, expr.Operator, right)
		}
		return fmt.Sprintf("%s %s %s", left, expr.Operator, right)
	default:
		return "/* unsupported expression */"
	}
}
