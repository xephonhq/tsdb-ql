package while

import "fmt"

// === Start of Boolean Expressions

//BoolValExp wrap boolean
type BoolValExp struct {
	b bool
}

//Eval for BoolValExp return boolean
func (v BoolValExp) Eval(s State) bool {
	return v.b
}

func (v BoolValExp) String() string {
	return fmt.Sprintf("%v", v.b)
}

// BoolEqExp represent  left equal to right
type BoolEqExp struct {
	l ArithExp
	r ArithExp
}

// BoolLesExp represent left strict less than right
type BoolLesExp struct {
	l ArithExp
	r ArithExp
}

// TODO: we have < but don't have >

//
type BoolNotExp struct {
	b BoolExp
}

//
type BoolAndExp struct {
	l BoolExp
	r BoolExp
}

//
type BoolOrExp struct {
	l BoolExp
	r BoolExp
}

// Eval for BoolEqExp evaluate left equal right
func (equ BoolEqExp) Eval(s State) bool {
	return equ.l.Eval(s) == equ.r.Eval(s)
}

func (equ BoolEqExp) String() string {
	return fmt.Sprintf("%s = %s", equ.l, equ.r)
}

// Eval for BoolLesExp evaluate left less than right
func (les BoolLesExp) Eval(s State) bool {
	return les.l.Eval(s) < les.r.Eval(s)
}

func (les BoolLesExp) String() string {
	return fmt.Sprintf("%s < %s", les.l, les.r)
}

//
func (not BoolNotExp) Eval(s State) bool {
	return !not.b.Eval(s)
}

func (not BoolNotExp) String() string {
	return fmt.Sprintf("!%s", not.b)
}

//
func (and BoolAndExp) Eval(s State) bool {
	return and.l.Eval(s) && and.r.Eval(s)
}

func (and BoolAndExp) String() string {
	return fmt.Sprintf("%s and %s", and.l, and.r)
}

func (or BoolOrExp) Eval(s State) bool {
	return or.l.Eval(s) || or.r.Eval(s)
}

func (or BoolOrExp) String() string {
	return fmt.Sprintf("%s or %s", or.l, or.r)
}

// === End of Boolean Expressions
