package semantic

import (
	"fmt"
	"tiny-basic/src/ast"
)

type SymbolEntry struct {
	Name  string
	Value ast.Expression
	Used  bool
}

type SymbolTable struct {
	variables map[string]*SymbolEntry
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{variables: make(map[string]*SymbolEntry)}
}

func (st *SymbolTable) DeclareVariable(name string, value ast.Expression) error {
	if _, exists := st.variables[name]; exists {
		return fmt.Errorf("variable '%s' is already declared", name)
	}
	st.variables[name] = &SymbolEntry{Name: name, Value: value, Used: false}
	return nil
}

func (st *SymbolTable) AssignVariable(name string, value ast.Expression) error {
	if entry, exists := st.variables[name]; exists {
		entry.Value = value
		return nil
	}
	return fmt.Errorf("variable '%s' not declared", name)
}

func (st *SymbolTable) GetVariable(name string) (ast.Expression, error) {
	entry, exists := st.variables[name]
	if exists {
		entry.Used = true
		return entry.Value, nil
	}
	return nil, fmt.Errorf("variable '%s' not declared", name)
}
