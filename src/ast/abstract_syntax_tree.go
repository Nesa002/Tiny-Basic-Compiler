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

type PrintStatement struct {
	Expression Expression
}

func (ps *PrintStatement) statementNode() {}

type LetStatement struct {
	Identifier Identifier
	Value      Expression
}

func (as *LetStatement) statementNode() {}

type AssignmentStatement struct {
	Identifier Identifier
	Value      Expression
}

func (as *AssignmentStatement) statementNode() {}

type IfStatement struct {
	Condition  Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (ifs *IfStatement) statementNode() {}

type WhileStatement struct {
	Condition Expression
	DoBranch  []Statement
}

func (ifs *WhileStatement) statementNode() {}

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
