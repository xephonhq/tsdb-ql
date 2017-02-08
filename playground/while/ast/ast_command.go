package ast

import "fmt"

// === start of Command Expression
type CommandSkipExp struct {
	//nothing
}

//Eval for CommandSkip Expression
func (skipExpr CommandSkipExp) Eval(s State) {
	//do nothing
}

func (skipExpr CommandSkipExp) EvalS(s State) CommandExp{
	//do nothing
	return CommandSkipExp{}
}

func (skipExpr CommandSkipExp) String() string {
	return "skip"
}

type CommandAssignExp struct {
	v VarExpr
	e ArithExp
}

func (assignExpr CommandAssignExp) Eval(s State) {
	s[assignExpr.v.Name()] = assignExpr.e.Eval(s)
}


func (assignExpr CommandAssignExp) EvalS(s State) CommandExp{
	s[assignExpr.v.Name()] = assignExpr.e.Eval(s)
	return CommandSkipExp{}
}


func (assignExpr CommandAssignExp) String() string {
	return fmt.Sprintf("%s := %s", assignExpr.v, assignExpr.e)
}

type CommandSeqExp struct {
	c1 CommandExp
	c2 CommandExp
}

func (seqExpr CommandSeqExp) Eval(s State) {
	seqExpr.c1.Eval(s) //s1
	seqExpr.c2.Eval(s) //s2
}

func (seqExpr CommandSeqExp) EvalS(s State ) CommandExp {
	_, ok := seqExpr.c1.(CommandSkipExp)
	if ok {
		return seqExpr.c2
	} else {
		return CommandSeqExp{c1: seqExpr.c1.EvalS(s), c2: seqExpr.c2}
	}
}

func (seqExpr CommandSeqExp) String() string {
	return fmt.Sprintf("%s;%s ", seqExpr.c1, seqExpr.c2)
}

type CommandIfExp struct {
	c1 CommandExp
	c2 CommandExp
	b  BoolExp
}

func (ifExpr CommandIfExp) Eval(s State) {
	if ifExpr.b.Eval(s) {
		ifExpr.c1.Eval(s)
	} else {
		ifExpr.c2.Eval(s)
	}
}

func (ifExpr CommandIfExp) EvalS(s State) CommandExp {
	if ifExpr.b.Eval(s) {
		return ifExpr.c1.EvalS(s)
	} else {
		return  ifExpr.c2.EvalS(s)
	}
}

func (ifExpr CommandIfExp) String() string {
	return fmt.Sprintf("if %s then %s else %s", ifExpr.b, ifExpr.c1, ifExpr.c2)
}

type CommandWhileExp struct {
	c CommandExp
	b BoolExp
	v VarExpr
}

func (whileExpr CommandWhileExp) Eval(s State) {
	if whileExpr.b.Eval(s) {
		whileExpr.c.Eval(s)
		whileExpr.Eval(s)
	} else {
		//skip
	}
}

func (whileExpr CommandWhileExp) EvalS(s State) CommandExp {
	if whileExpr.b.Eval(s) {
		return whileExpr.c.EvalS(s)
		whileExpr.EvalS(s)
	} else {
		return CommandSkipExp{}
	}
}

func (whileExpr CommandWhileExp) String() string {
	return fmt.Sprintf("while %s then %s ", whileExpr.b, whileExpr.c)
}
