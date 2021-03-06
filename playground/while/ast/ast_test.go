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

func assertEvalAll(exp CommandExp, postState State, t *testing.T) {
	s := make(map[string]int)
	EvalAll(exp,s)

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
		//fmt.Printf("%s => %v\n", exp, postState)
	}
}

// TODO： test -， test *
func TestArithSumExp_Eval(t *testing.T) {
	var asst = func(exp ArithExp, v int) {
		assertArithExp(exp, v, t)
	}
	asst(ArithSumExp{l: ArithValExp{num: 1}, r: ArithValExp{num: 2}}, 3)
}

func TestArithSubExp_Eval(t *testing.T) {
	var asst = func(exp ArithExp, v int) {
		assertArithExp(exp, v, t)
	}
	asst(ArithSubExp{l: ArithValExp{num: 1}, r: ArithValExp{num: 2}}, -1)
}

func TestArithMulExp_Eval(t *testing.T) {
	var asst = func(exp ArithExp, v int) {
		assertArithExp(exp, v, t)
	}
	asst(ArithMulExp{l: ArithValExp{num: 5}, r: ArithValExp{num: 6}}, 30)
}

func TestArithDivExp_Eval(t *testing.T) {
	var asst = func(exp ArithExp, v int) {
		assertArithExp(exp, v, t)
	}
	asst(ArithDivExp{l: ArithValExp{num: 22}, r: ArithValExp{num: 2}}, 11)
}

// TODO: or, not, not eq, gt, les
func TestBoolAndExp_Eval(t *testing.T) {
	var asst = func(exp BoolExp, b bool) {
		assertBoolExp(exp, b, t)
	}
	asst(BoolAndExp{l: BoolValExp{b: true}, r: BoolValExp{b: false}}, false)
}

func TestBoolOrExp_Eval(t *testing.T) {
	var asst = func(exp BoolExp, b bool) {
		assertBoolExp(exp, b, t)
	}
	asst(BoolOrExp{l: BoolValExp{b: true}, r: BoolValExp{b: false}}, true)
}

func TestBoolNotExp_Eval(t *testing.T) {
	var asst = func(exp BoolExp, b bool) {
		assertBoolExp(exp, b, t)
	}
	asst(BoolNotExp{b: BoolValExp{b: false}}, true)
}

// TODO: if, seq, assign Though, this test should cover most of them
/*func TestCommandAssignExp_Eval(t *testing.T) {
	var asst = func(exp CommandExp, s State) {
		assertCommandExp(exp, s, t)
	}

	var exp = CommandAssignExp{v: VarExpr{name: "x"}, e: ArithValExp{num: 1}}
	asst(exp, map[string]int{"x": 1})
}*/

func TestCommandWhileExp_Eval(t *testing.T) {
	var asst = func(exp CommandExp, s State) {
		assertCommandExp(exp, s, t)
	}

	var exp = CommandSeqExp{c1: CommandAssignExp{v: VarExpr{name: "x"}, e: ArithValExp{num: 1}},
		c2: CommandWhileExp{b: BoolLesExp{l: VarExpr{name: "x"}, r: ArithValExp{num: 200}},
			c: CommandAssignExp{v: VarExpr{name: "x"},
				e: ArithSumExp{l: VarExpr{name: "x"},
					r:        ArithValExp{num: 1}}}}}
	asst(exp, map[string]int{"x": 200})
}

func TestCommandIfExp_Eval(t *testing.T) {
	var asst = func(exp CommandExp, s State) {
		assertCommandExp(exp, s, t)
	}

	var exp = CommandSeqExp{c1: CommandAssignExp{v: VarExpr{name: "x"}, e: ArithValExp{num: 1}},
		c2: CommandIfExp{b: BoolLesExp{l: VarExpr{name: "x"}, r: ArithValExp{6} },
			c1: CommandAssignExp{v: VarExpr{name: "x"},
				e: ArithSumExp{l: VarExpr{name: "x"},
					r:        ArithValExp{num: 1}}},
			c2: CommandSkipExp{}}}

	asst(exp, map[string]int{"x": 2})
}

// TODO: small step test


func TestCommandWhileExp_EvalS(t *testing.T) {
	var asst = func(exp CommandExp, s State) {
		assertEvalAll(exp, s, t)
	}

	var exp = CommandSeqExp{c1: CommandAssignExp{v: VarExpr{name: "x"}, e: ArithValExp{num: 1}},
		c2: CommandWhileExp{b: BoolLesExp{l: VarExpr{name: "x"}, r: ArithValExp{num: 6}},
			c: CommandAssignExp{v: VarExpr{name: "x"},
				e: ArithSumExp{l: VarExpr{name: "x"},
					r:        ArithValExp{num: 1}}}}}

	asst(exp, map[string]int{"x": 6})
}

func TestCommandIfExp_EvalS(t *testing.T) {
	var asst = func(exp CommandExp, s State) {
		assertEvalAll(exp, s, t)
	}

	var exp = CommandSeqExp{c1: CommandAssignExp{v: VarExpr{name: "x"}, e: ArithValExp{num: 1}},
		c2: CommandIfExp{b: BoolLesExp{l: VarExpr{name: "x"}, r: ArithValExp{6} },
			c1: CommandAssignExp{v: VarExpr{name: "x"},
				e: ArithSumExp{l: VarExpr{name: "x"},
					r:        ArithValExp{num: 1}}},
			c2: CommandSkipExp{}}}

	asst(exp, map[string]int{"x": 2})
}
