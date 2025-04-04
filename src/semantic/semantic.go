package semantic

import (
	"fmt"
	"tiny-basic/src/ast"
)

type SemanticAnalyzer struct {
	symbolTable *SymbolTable
}

func NewSemanticAnalyzer() *SemanticAnalyzer {
	return &SemanticAnalyzer{symbolTable: NewSymbolTable()}
}

func (sa *SemanticAnalyzer) Analyze(program *ast.Program) error {
	for _, stmt := range program.Statements {
		if err := sa.analyzeStatement(stmt); err != nil {
			return err
		}
	}

	return nil
}

func (sa *SemanticAnalyzer) CheckUnusedVariables() []string {
	var warnings []string

	for _, entry := range sa.symbolTable.variables {
		if entry.Used {
			continue
		}
		warnings = append(warnings, fmt.Sprintf("Warning: Variable '%s' is declared but never used", entry.Name))
	}
	return warnings
}

func (sa *SemanticAnalyzer) analyzeStatement(stmt ast.Statement) error {
	switch stmt := stmt.(type) {
	case *ast.LetStatement:
		return sa.analyzeLetStatement(stmt)
	case *ast.AssignmentStatement:
		return sa.analyzeAssignmentStatement(stmt)
	case *ast.PrintStatement:
		return sa.analyzePrintStatement(stmt)
	case *ast.IfStatement:
		return sa.analyzeIfStatement(stmt)
	case *ast.WhileStatement:
		return sa.analyzeWhileStatement(stmt)
	case *ast.EndStatement, *ast.CommentStatement:
		return nil
	default:
		return fmt.Errorf("unknown statement type")
	}
}

func (sa *SemanticAnalyzer) analyzeLetStatement(stmt *ast.LetStatement) error {
	if err := sa.symbolTable.DeclareVariable(stmt.Identifier.Name, stmt.Value); err != nil {
		return err
	}
	return sa.analyzeExpression(stmt.Value)
}

func (sa *SemanticAnalyzer) analyzeAssignmentStatement(stmt *ast.AssignmentStatement) error {
	return sa.symbolTable.AssignVariable(stmt.Identifier.Name, stmt.Value)
}

func (sa *SemanticAnalyzer) analyzePrintStatement(stmt *ast.PrintStatement) error {
	return sa.analyzeExpression(stmt.Expression)
}

func (sa *SemanticAnalyzer) analyzeIfStatement(stmt *ast.IfStatement) error {
	if err := sa.analyzeExpression(stmt.Condition); err != nil {
		return err
	}

	if !sa.isBooleanExpression(stmt.Condition) {
		return fmt.Errorf("condition in IF statement must be a comparison (==, <, >), got: %T", stmt.Condition)
	}

	if err := sa.analyzeStatement(stmt.ThenBranch); err != nil {
		return err
	}

	if stmt.ElseBranch != nil {
		if err := sa.analyzeStatement(stmt.ElseBranch); err != nil {
			return err
		}
	}

	return nil
}

func (sa *SemanticAnalyzer) analyzeWhileStatement(stmt *ast.WhileStatement) error {
	if err := sa.analyzeExpression(stmt.Condition); err != nil {
		return err
	}

	if !sa.isBooleanExpression(stmt.Condition) {
		return fmt.Errorf("condition in WHILE statement must be a comparison (==, <, >), got: %T", stmt.Condition)
	}

	for _, statement := range stmt.DoBranch {
		if err := sa.analyzeStatement(statement); err != nil {
			return err
		}
	}
	return nil
}

func (sa *SemanticAnalyzer) analyzeExpression(expr ast.Expression) error {
	switch expr := expr.(type) {
	case *ast.IntegerLiteral, *ast.FloatLiteral:
		return nil
	case *ast.Identifier:
		_, err := sa.symbolTable.GetVariable(expr.Name)
		return err
	case *ast.BinaryExpression:
		if err := sa.analyzeExpression(expr.Left); err != nil {
			return err
		}
		if err := sa.analyzeExpression(expr.Right); err != nil {
			return err
		}
	}
	return nil
}

func (sa *SemanticAnalyzer) isBooleanExpression(expr ast.Expression) bool {
	binExpr, ok := expr.(*ast.BinaryExpression)
	if !ok {
		return false
	}

	switch binExpr.Operator {
	case "==", "<", ">":
		return true
	default:
		return false
	}
}
