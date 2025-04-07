package main

import (
	"fmt"
	"os"
	"tiny-basic/src/codegen"
	"tiny-basic/src/optimizer"
	"tiny-basic/src/parser"
	"tiny-basic/src/semantic"
	"tiny-basic/src/tokenizer"
)

func main() {
	inputFile := "input.tb"
	outputFile := "output.js"

	sourceCode, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	tokens, err := tokenizer.Tokenize(string(sourceCode))
	if err != nil {
		fmt.Println(err)
		return
	}

	p := parser.NewParser(tokens)

	program := p.ParseProgram()
	program = optimizer.Optimize(program)

	sa := semantic.NewSemanticAnalyzer()
	if err := sa.Analyze(program); err != nil {
		fmt.Println("Error during semantic analysis:", err)
		return
	}
	warnings := sa.CheckUnusedVariables()
	for _, warning := range warnings {
		fmt.Println(warning)
	}

	cg := codegen.NewCodeGenerator()
	jsCode := cg.Generate(program)

	// Write the generated JavaScript code to a file
	err = os.WriteFile(outputFile, []byte(jsCode), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	fmt.Println("Compilation successful! JavaScript output saved to", outputFile)

}
