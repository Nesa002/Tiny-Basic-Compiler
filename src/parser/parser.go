package parser

import (
	"fmt"
	"strconv"
	"tiny-basic/src/ast"
	"tiny-basic/src/tokenizer"
)

type Parser struct {
	tokens  []tokenizer.Token
	current int
}

func NewParser(tokens []tokenizer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.current < len(p.tokens)-1 {
		statement := p.parseStatement()
		program.Statements = append(program.Statements, statement)
		if statement == nil {
			return program
		}
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.tokens[p.current].Type {
	case tokenizer.TOKEN_LET:
		return p.parseLetStatement()
	case tokenizer.TOKEN_IF:
		return p.parseIfStatement()
	case tokenizer.TOKEN_WHILE:
		return p.parseWhileStatement()
	case tokenizer.TOKEN_PRINT:
		return p.parsePrintStatement()
	case tokenizer.TOKEN_COMMENT:
		return p.parseCommentStatement()
	case tokenizer.TOKEN_END:
		return p.parseEndStatement()
	case tokenizer.TOKEN_EOF:
		return nil
	}

	if p.peek().Type == tokenizer.TOKEN_IDENTIFIER {
		return p.parseAssignmentStatement()
	}

	p.parseError("Unknown Statement")
	return nil
}

func (p *Parser) parseLetStatement() ast.Statement {
	p.consume(tokenizer.TOKEN_LET, "Expected LET keyword")
	varName := p.consume(tokenizer.TOKEN_IDENTIFIER, "Expected an identifier")
	p.consume(tokenizer.TOKEN_EQUALS, "Expected '=' operator for assignment")
	value := p.parseExpression()

	return &ast.LetStatement{
		Identifier: ast.Identifier{Name: varName.Value},
		Value:      value,
	}
}

func (p *Parser) parseAssignmentStatement() ast.Statement {
	varName := p.consume(tokenizer.TOKEN_IDENTIFIER, "Expected variable name (identifier)")
	p.consume(tokenizer.TOKEN_EQUALS, "Expected '=' operator for assignment")

	value := p.parseExpression()

	return &ast.AssignmentStatement{
		Identifier: ast.Identifier{Name: varName.Value},
		Value:      value,
	}
}

func (p *Parser) parseIfStatement() ast.Statement {
	p.consume(tokenizer.TOKEN_IF, "Expected IF keyword")
	condition := p.parseExpression()
	p.consume(tokenizer.TOKEN_THEN, "Exprected THEN keyword after condition")
	thenBranch := p.parseStatement()

	var elseBranch ast.Statement = nil
	if p.peek().Type == tokenizer.TOKEN_ELSE {
		p.consume(tokenizer.TOKEN_ELSE, "Expected ELSE keyword")
		elseBranch = p.parseStatement()
	}

	return &ast.IfStatement{
		Condition:  condition,
		ThenBranch: thenBranch,
		ElseBranch: elseBranch,
	}
}

func (p *Parser) parseWhileStatement() ast.Statement {
	p.consume(tokenizer.TOKEN_WHILE, "Expected WHILE keyword")
	condition := p.parseExpression()
	p.consume(tokenizer.TOKEN_DO, "Exprected DO keyword after condition")

	doBranch := []ast.Statement{}

	for p.peek().Type != tokenizer.TOKEN_STOP {
		doBranch = append(doBranch, p.parseStatement())
	}
	p.consume(tokenizer.TOKEN_STOP, "Exprected STOP keyword after condition")

	return &ast.WhileStatement{
		Condition: condition,
		DoBranch:  doBranch,
	}
}

func (p *Parser) parsePrintStatement() ast.Statement {
	p.consume(tokenizer.TOKEN_PRINT, "Exprected PRINT keyword")
	expression := p.parseExpression()

	return &ast.PrintStatement{
		Expression: expression,
	}
}

func (p *Parser) parseCommentStatement() ast.Statement {
	text := p.consume(tokenizer.TOKEN_COMMENT, "Expected comment")

	return &ast.CommentStatement{
		Text: text.Value,
	}
}

func (p *Parser) parseEndStatement() ast.Statement {
	p.consume(tokenizer.TOKEN_END, "Expected END keyword")

	return &ast.EndStatement{}
}

func (p *Parser) parseExpression() ast.Expression {
	left := p.parseTerm()

	for p.peek().Type == tokenizer.TOKEN_REL_OP {
		operator := p.consume(p.peek().Type, "Expected relational operator.").Value
		right := p.parseTerm()

		left = &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left
}

func (p *Parser) parseTerm() ast.Expression {
	left := p.parseFactor()

	for p.peek().Type == tokenizer.TOKEN_ADD_SUB {
		operator := p.consume(tokenizer.TOKEN_ADD_SUB, "Expected + or -").Value
		right := p.parseFactor()
		left = &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left
}

func (p *Parser) parseFactor() ast.Expression {
	left := p.parsePrimaryExpression()

	for p.peek().Type == tokenizer.TOKEN_MUL_DIV {
		operator := p.consume(tokenizer.TOKEN_MUL_DIV, "Expected * or /").Value
		right := p.parsePrimaryExpression()
		left = &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}

	return left
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	if p.match(tokenizer.TOKEN_INTEGER) {
		return &ast.IntegerLiteral{
			Value: atoi(p.previous().Value),
		}
	}
	if p.match(tokenizer.TOKEN_FLOAT) {
		return &ast.FloatLiteral{
			Value: atof(p.previous().Value),
		}
	}
	if p.match(tokenizer.TOKEN_IDENTIFIER) {
		return &ast.Identifier{
			Name: p.previous().Value,
		}
	}
	if p.match(tokenizer.TOKEN_LEFT_PAREN) {
		expression := p.parseExpression()
		p.consume(tokenizer.TOKEN_RIGHT_PAREN, "Expected closing parenthesis.")
		return expression
	}

	p.parseError("Exprected expression")
	return nil
}

func (p *Parser) peek() tokenizer.Token {
	return p.tokens[p.current]
}

func (p *Parser) consume(expected tokenizer.TokenType, msg string) tokenizer.Token {
	if p.match(expected) {
		return p.previous()
	}
	p.parseError(msg)
	return tokenizer.Token{}
}

func (p *Parser) match(tokenType tokenizer.TokenType) bool {
	if p.current < len(p.tokens) && p.tokens[p.current].Type == tokenType {
		p.current++
		return true
	}
	return false
}

func (p *Parser) previous() tokenizer.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) parseError(msg string) {
	panic(fmt.Sprintf("Parse error at line %d, token {%s: %s}: %s", p.tokens[p.current].Line, p.tokens[p.current].Type, p.tokens[p.current].Value, msg))
}

func atof(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic("Invalid number format")
	}
	return val
}

func atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic("Invalid integer format")
	}
	return val
}
