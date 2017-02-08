package ast

import (
	"fmt"
	"testing"
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

func assertCommandExp(exp CommandExp, postState State, t *testing.T) {
	s := make(map[string]int)
	exp.Eval(s)
	// compare state
	fail := false
	for k, v := range postState {
		if s[k] != v {
			fmt.Printf("expect %s to be %d but got %d\n", k, v, s[k])
			fail = true
		}
	}
	if fail {
		t.Fail()
	} else {
		fmt.Printf("%s => %v\n", exp, postState)
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

// TODO: if, seq, assign Though, this test should cover most of them
func TestCommandWhileExp_Eval(t *testing.T) {
	var asst = func(exp CommandExp, s State) {
		assertCommandExp(exp, s, t)
	}

	var exp = CommandSeqExp{c1: CommandAssignExp{v: VarExpr{name: "x"}, e: ArithValExp{num: 1}},
		c2: CommandWhileExp{b: BoolLesExp{l: VarExpr{name: "x"}, r: ArithValExp{num: 100}},
			c: CommandAssignExp{v: VarExpr{name: "x"}, e: ArithSumExp{l: VarExpr{name: "x"}, r: ArithValExp{num: 100}}}},
	}

	asst(exp, map[string]int{"x": 101})
}

// TODO: small step test
