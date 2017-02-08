package while

import "fmt"

// ArithValExp wrap raw integers
type ArithValExp struct {
	num int
}

// Eval for ArithValExp returns number
func (v ArithValExp) Eval(s State) int {
	return v.num
}

func (v ArithValExp) String() string {
	return fmt.Sprintf("%d", v.num)
}

// ArithSumExp represent sum
type ArithSumExp struct {
	l ArithExp
	r ArithExp
}

// ArithSubExp represent Subtraction
type ArithSubExp struct {
	l ArithExp
	r ArithExp
}

// ArithMulExp represent multiplication
type ArithMulExp struct {
	l ArithExp
	r ArithExp
}

// ArithDivExp represent division
type ArithDivExp struct {
	l ArithExp
	r ArithExp
}

// Eval for ArithSumExp evaluate left and right and add them up
func (sum ArithSumExp) Eval(s State) int {
	return sum.l.Eval(s) + sum.r.Eval(s)
}

func (sum ArithSumExp) String() string {
	return fmt.Sprintf("(%s + %s)", sum.l, sum.r)
}

// Eval for ArithSubExp evaluate left minus right
func (sub ArithSubExp) Eval(s State) int {
	return sub.l.Eval(s) - sub.r.Eval(s)
}

func (sub ArithSubExp) String() string {
	return fmt.Sprintf("(%s - %s)", sub.l, sub.r)
}

// Eval for ArithMulxp evaluate left product right
func (mul ArithMulExp) Eval(s State) int {
	return mul.l.Eval(s) * mul.r.Eval(s)
}

func (mul ArithMulExp) String() string {
	return fmt.Sprintf("%s * %s", mul.l, mul.r)
}

// Eval for ArithDivExp evaluate left divided by right
// TODO: handle divide by zero
func (div ArithDivExp) Eval(s State) int {
	return div.l.Eval(s) / div.r.Eval(s)
}

func (div ArithDivExp) String() string {
	return fmt.Sprintf("%s / %s", div.l, div.r)
}

// === End of Arithmetic Expressions
