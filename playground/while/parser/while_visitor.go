// Generated from while.g4 by ANTLR 4.6.

package parser // while

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by whileParser.
type whileVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by whileParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by whileParser#div.
	VisitDiv(ctx *DivContext) interface{}

	// Visit a parse tree produced by whileParser#sub.
	VisitSub(ctx *SubContext) interface{}

	// Visit a parse tree produced by whileParser#mul.
	VisitMul(ctx *MulContext) interface{}

	// Visit a parse tree produced by whileParser#var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by whileParser#num.
	VisitNum(ctx *NumContext) interface{}

	// Visit a parse tree produced by whileParser#sum.
	VisitSum(ctx *SumContext) interface{}

	// Visit a parse tree produced by whileParser#les.
	VisitLes(ctx *LesContext) interface{}

	// Visit a parse tree produced by whileParser#gt.
	VisitGt(ctx *GtContext) interface{}

	// Visit a parse tree produced by whileParser#eq.
	VisitEq(ctx *EqContext) interface{}

	// Visit a parse tree produced by whileParser#neq.
	VisitNeq(ctx *NeqContext) interface{}

	// Visit a parse tree produced by whileParser#bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by whileParser#skip.
	VisitSkip(ctx *SkipContext) interface{}

	// Visit a parse tree produced by whileParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by whileParser#if.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by whileParser#while.
	VisitWhile(ctx *WhileContext) interface{}
}
