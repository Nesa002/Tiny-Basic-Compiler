package ast

type Node interface{}

type Program struct {
	Statements []Statement
}

// Statements
type Statement interface {
	Node
	statementNode()
}

// This is used for standalone Expressions like "Hello" in PRINT "Hello"
type ExpressionStatement struct {
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

type PrintStatement struct {
	Expression Expression
}

func (ps *PrintStatement) statementNode() {}

type AssignmentStatement struct {
	Name  string
	Value Expression
}

func (as *AssignmentStatement) statementNode() {}

type IfStatement struct {
	Condition  Expression
	ThenBranch Statement
}

func (ifs *IfStatement) statementNode() {}

type CommentStatement struct {
	Text string
}

func (cs *CommentStatement) statementNode() {}

type EndStatement struct{}

func (es *EndStatement) statementNode() {}

// Expressions
type Expression interface {
	Node
	expressionNode()
}

type IntegerLiteral struct {
	Value int
}

func (il *IntegerLiteral) expressionNode() {}

type FloatLiteral struct {
	Value float64
}

func (fl *FloatLiteral) expressionNode() {}

type StringLiteral struct {
	Value string
}

func (sl *StringLiteral) expressionNode() {}

type Identifier struct {
	Name string
}

func (id *Identifier) expressionNode() {}

type BinaryExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) expressionNode() {}
