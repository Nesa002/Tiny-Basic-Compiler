package main

import (
	"fmt"
	"os"
	"tiny-basic/src/ast"
	"tiny-basic/src/parser"
	"tiny-basic/src/tokenizer"
)

func main() {
	fileName := "example.tb"
	sourceCode, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	tokens, err := tokenizer.Tokenize(string(sourceCode))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, token := range tokens {
		fmt.Printf("Token(Type: %s, Value: %s)\n", token.Type, token.Value)
	}

	p := parser.NewParser(tokens)

	program := p.ParseProgram()
	fmt.Println("Parsed Expression AST:")
	printAST(*program, 0)
}

// Helper function to pretty print the AST for a program
func printAST(program ast.Program, indent int) {
	// Indentation for better readability
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}

	// Assuming Program is a collection of statements
	for _, stmt := range program.Statements {
		printStatement(stmt, indent)
	}
}

// Helper function to print each statement
func printStatement(stmt ast.Statement, indent int) {
	// Handle different statement types (e.g., Assignment, Expression)

	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}

	switch s := stmt.(type) {
	case *ast.AssignmentStatement:
		fmt.Printf("%sAssignmentStatement:\n", indentStr)
		fmt.Printf("%s  Name: %s\n", indentStr, s.Name)
		fmt.Printf("%s  Value:\n", indentStr)
		printExpression(s.Value, indent+2)
	case *ast.PrintStatement:
		fmt.Printf("%sPrintStatement:\n", indentStr)
		fmt.Printf("%s  Expression:\n", indentStr)
		printExpression(s.Expression, indent+2)
	default:
		fmt.Printf("%sUnknown statement type\n", indentStr)
	}
}

// Helper function to print expressions
func printExpression(expr ast.Expression, indent int) {
	// Indentation for better readability
	indentStr := ""
	for i := 0; i < indent; i++ {
		indentStr += "  "
	}

	switch e := expr.(type) {
	case *ast.IntegerLiteral:
		fmt.Printf("%sIntegerLiteral: %d\n", indentStr, e.Value)
	case *ast.FloatLiteral:
		fmt.Printf("%sFloatLiteral: %f\n", indentStr, e.Value)
	case *ast.Identifier:
		fmt.Printf("%sIdentifier: %s\n", indentStr, e.Name)
	case *ast.BinaryExpression:
		fmt.Printf("%sBinaryExpression:\n", indentStr)
		fmt.Printf("%s  Left:\n", indentStr)
		printExpression(e.Left, indent+2)
		fmt.Printf("%s  Operator: %s\n", indentStr, e.Operator)
		fmt.Printf("%s  Right:\n", indentStr)
		printExpression(e.Right, indent+2)
	default:
		fmt.Printf("%sUnknown expression type\n", indentStr)
	}
}
