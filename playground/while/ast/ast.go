package ast

import "fmt"

// State is just simple map
type State map[string]int

// Exp is the base interface for Expression
type Exp interface {
	String() string
}

// ArithExp is the required interface for Arithmetic Expression
type ArithExp interface {
	Eval(s State) int
	Exp
}

// BoolExp is the required interface for Boolean Expression
type BoolExp interface {
	Eval(s State) bool
	Exp
}

// CommandExp is the required interface for Command Expression
type CommandExp interface {
	Eval(s State)
	// TODO: Add EvalS for small step
	EvalS(s State) CommandExp
	Exp
}

// === Start of Arithmetic Expressions

// === Start of Variable Expression

// VarExpr is ?
type VarExpr struct {
	name string
	Exp
}

func (v VarExpr) Name() string {
	return v.name
}

func (v VarExpr) Eval(s State) int {
	val, ok := s[v.name]
	if !ok {
		panic(fmt.Sprintf("%s is undefined variable", v.name))
	}
	return val
}

func (v VarExpr) String() string {
	return v.name
}

// === End of Variable Expression