// Generated from while.g4 by ANTLR 4.6.

package parser // while

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasewhileListener is a complete listener for a parse tree produced by whileParser.
type BasewhileListener struct{}

var _ whileListener = &BasewhileListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasewhileListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasewhileListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasewhileListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasewhileListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasewhileListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasewhileListener) ExitProg(ctx *ProgContext) {}

// EnterDiv is called when production div is entered.
func (s *BasewhileListener) EnterDiv(ctx *DivContext) {}

// ExitDiv is called when production div is exited.
func (s *BasewhileListener) ExitDiv(ctx *DivContext) {}

// EnterSub is called when production sub is entered.
func (s *BasewhileListener) EnterSub(ctx *SubContext) {}

// ExitSub is called when production sub is exited.
func (s *BasewhileListener) ExitSub(ctx *SubContext) {}

// EnterMul is called when production mul is entered.
func (s *BasewhileListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production mul is exited.
func (s *BasewhileListener) ExitMul(ctx *MulContext) {}

// EnterVar is called when production var is entered.
func (s *BasewhileListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BasewhileListener) ExitVar(ctx *VarContext) {}

// EnterNum is called when production num is entered.
func (s *BasewhileListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BasewhileListener) ExitNum(ctx *NumContext) {}

// EnterSum is called when production sum is entered.
func (s *BasewhileListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production sum is exited.
func (s *BasewhileListener) ExitSum(ctx *SumContext) {}

// EnterLes is called when production les is entered.
func (s *BasewhileListener) EnterLes(ctx *LesContext) {}

// ExitLes is called when production les is exited.
func (s *BasewhileListener) ExitLes(ctx *LesContext) {}

// EnterGt is called when production gt is entered.
func (s *BasewhileListener) EnterGt(ctx *GtContext) {}

// ExitGt is called when production gt is exited.
func (s *BasewhileListener) ExitGt(ctx *GtContext) {}

// EnterEq is called when production eq is entered.
func (s *BasewhileListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production eq is exited.
func (s *BasewhileListener) ExitEq(ctx *EqContext) {}

// EnterNeq is called when production neq is entered.
func (s *BasewhileListener) EnterNeq(ctx *NeqContext) {}

// ExitNeq is called when production neq is exited.
func (s *BasewhileListener) ExitNeq(ctx *NeqContext) {}

// EnterBool is called when production bool is entered.
func (s *BasewhileListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BasewhileListener) ExitBool(ctx *BoolContext) {}

// EnterSkip is called when production skip is entered.
func (s *BasewhileListener) EnterSkip(ctx *SkipContext) {}

// ExitSkip is called when production skip is exited.
func (s *BasewhileListener) ExitSkip(ctx *SkipContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasewhileListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasewhileListener) ExitAssign(ctx *AssignContext) {}

// EnterIf is called when production if is entered.
func (s *BasewhileListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BasewhileListener) ExitIf(ctx *IfContext) {}

// EnterWhile is called when production while is entered.
func (s *BasewhileListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BasewhileListener) ExitWhile(ctx *WhileContext) {}
