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

	expr := p.ParseExpression()
	fmt.Println("Parsed Expression AST:")
	printAST(expr, 0)
}

// Helper function to pretty print the AST (you can adjust this function based on your AST structure)
func printAST(expr ast.Expression, indent int) {
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
		printAST(e.Left, indent+2)
		fmt.Printf("%s  Operator: %s\n", indentStr, e.Operator)
		fmt.Printf("%s  Right:\n", indentStr)
		printAST(e.Right, indent+2)
	default:
		fmt.Printf("%sUnknown expression type\n", indentStr)
	}
}
