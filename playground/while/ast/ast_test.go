package ast

import (
	"testing"
	"fmt"
)

func assertArithExp(exp ArithExp, v int, t *testing.T) {
	s := make(map[string]int)
	r := exp.Eval(s)
	if r != v {
		fmt.Printf("Fail to evaluate %s\n", exp)
		fmt.Printf("expect %d but got %d\n", r, v)
		t.Fail()
	} else {
		fmt.Printf("%s => %d\n", exp, r)
	}
}

func assertBoolExp(exp BoolExp, b bool, t *testing.T) {
	s := make(map[string]int)
	r := exp.Eval(s)
	if r != b {
		fmt.Printf("Fail to evaluate %s\n", exp)
		fmt.Printf("expect %v but got %v\n", r, b)
		t.Fail()
	} else {
		fmt.Printf("%s => %v\n", exp, r)
	}
}

// TODO： test -， test *
func TestArithSumExp_Eval(t *testing.T) {
	var asst = func(exp ArithExp, v int) {
		assertArithExp(exp, v, t)
	}
	asst(ArithSumExp{l: ArithValExp{num: 1}, r: ArithValExp{num: 2}}, 3)
}


// TODO: or, not, not eq, gt, les
func TestBoolAndExp_Eval(t *testing.T) {
	var asst = func(exp BoolExp, b bool) {
		assertBoolExp(exp, b, t)
	}
	asst(BoolAndExp{l: BoolValExp{b: true}, r: BoolValExp{b: false}}, false)
}

