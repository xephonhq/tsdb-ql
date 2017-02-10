// Generated from while.g4 by ANTLR 4.6.

package parser // while

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasewhileVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasewhileVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitDiv(ctx *DivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitSub(ctx *SubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitMul(ctx *MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitSum(ctx *SumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitLes(ctx *LesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitGt(ctx *GtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitEq(ctx *EqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitNeq(ctx *NeqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitSkip(ctx *SkipContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitIf(ctx *IfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasewhileVisitor) VisitWhile(ctx *WhileContext) interface{} {
	return v.VisitChildren(ctx)
}
