// Generated from while.g4 by ANTLR 4.6.

package parser // while

import "github.com/antlr/antlr4/runtime/Go/antlr"

// whileListener is a complete listener for a parse tree produced by whileParser.
type whileListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterDiv is called when entering the div production.
	EnterDiv(c *DivContext)

	// EnterSub is called when entering the sub production.
	EnterSub(c *SubContext)

	// EnterMul is called when entering the mul production.
	EnterMul(c *MulContext)

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterSum is called when entering the sum production.
	EnterSum(c *SumContext)

	// EnterLes is called when entering the les production.
	EnterLes(c *LesContext)

	// EnterGt is called when entering the gt production.
	EnterGt(c *GtContext)

	// EnterEq is called when entering the eq production.
	EnterEq(c *EqContext)

	// EnterNeq is called when entering the neq production.
	EnterNeq(c *NeqContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterSkip is called when entering the skip production.
	EnterSkip(c *SkipContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterWhile is called when entering the while production.
	EnterWhile(c *WhileContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitDiv is called when exiting the div production.
	ExitDiv(c *DivContext)

	// ExitSub is called when exiting the sub production.
	ExitSub(c *SubContext)

	// ExitMul is called when exiting the mul production.
	ExitMul(c *MulContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitSum is called when exiting the sum production.
	ExitSum(c *SumContext)

	// ExitLes is called when exiting the les production.
	ExitLes(c *LesContext)

	// ExitGt is called when exiting the gt production.
	ExitGt(c *GtContext)

	// ExitEq is called when exiting the eq production.
	ExitEq(c *EqContext)

	// ExitNeq is called when exiting the neq production.
	ExitNeq(c *NeqContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitSkip is called when exiting the skip production.
	ExitSkip(c *SkipContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitWhile is called when exiting the while production.
	ExitWhile(c *WhileContext)
}
