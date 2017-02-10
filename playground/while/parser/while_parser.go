// Generated from while.g4 by ANTLR 4.6.

package parser // while

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 1072, 54993, 33286, 44333, 17431, 44785, 36224, 43741, 3, 25, 96, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 6, 2, 12, 10, 2, 13,
	2, 14, 2, 13, 3, 3, 3, 3, 3, 3, 5, 3, 19, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 33, 10, 3, 12, 3,
	14, 3, 36, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 55, 10, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5,
	69, 10, 5, 13, 5, 14, 5, 70, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 77, 10, 5, 13,
	5, 14, 5, 78, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 88, 10, 5,
	13, 5, 14, 5, 89, 3, 5, 3, 5, 5, 5, 94, 10, 5, 3, 5, 2, 3, 4, 6, 2, 4,
	6, 8, 2, 2, 107, 2, 11, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 54, 3, 2, 2,
	2, 8, 93, 3, 2, 2, 2, 10, 12, 5, 8, 5, 2, 11, 10, 3, 2, 2, 2, 12, 13, 3,
	2, 2, 2, 13, 11, 3, 2, 2, 2, 13, 14, 3, 2, 2, 2, 14, 3, 3, 2, 2, 2, 15,
	16, 8, 3, 1, 2, 16, 19, 7, 22, 2, 2, 17, 19, 7, 21, 2, 2, 18, 15, 3, 2,
	2, 2, 18, 17, 3, 2, 2, 2, 19, 34, 3, 2, 2, 2, 20, 21, 12, 8, 2, 2, 21,
	22, 7, 3, 2, 2, 22, 33, 5, 4, 3, 9, 23, 24, 12, 7, 2, 2, 24, 25, 7, 4,
	2, 2, 25, 33, 5, 4, 3, 8, 26, 27, 12, 6, 2, 2, 27, 28, 7, 5, 2, 2, 28,
	33, 5, 4, 3, 7, 29, 30, 12, 5, 2, 2, 30, 31, 7, 6, 2, 2, 31, 33, 5, 4,
	3, 6, 32, 20, 3, 2, 2, 2, 32, 23, 3, 2, 2, 2, 32, 26, 3, 2, 2, 2, 32, 29,
	3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2,
	35, 5, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 38, 5, 4, 3, 2, 38, 39, 7, 7,
	2, 2, 39, 40, 5, 4, 3, 2, 40, 55, 3, 2, 2, 2, 41, 42, 5, 4, 3, 2, 42, 43,
	7, 8, 2, 2, 43, 44, 5, 4, 3, 2, 44, 55, 3, 2, 2, 2, 45, 46, 5, 4, 3, 2,
	46, 47, 7, 9, 2, 2, 47, 48, 5, 4, 3, 2, 48, 55, 3, 2, 2, 2, 49, 50, 5,
	4, 3, 2, 50, 51, 7, 10, 2, 2, 51, 52, 5, 4, 3, 2, 52, 55, 3, 2, 2, 2, 53,
	55, 7, 23, 2, 2, 54, 37, 3, 2, 2, 2, 54, 41, 3, 2, 2, 2, 54, 45, 3, 2,
	2, 2, 54, 49, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 7, 3, 2, 2, 2, 56, 57,
	7, 11, 2, 2, 57, 94, 7, 12, 2, 2, 58, 59, 7, 21, 2, 2, 59, 60, 7, 13, 2,
	2, 60, 61, 5, 4, 3, 2, 61, 62, 7, 12, 2, 2, 62, 94, 3, 2, 2, 2, 63, 64,
	7, 14, 2, 2, 64, 65, 5, 6, 4, 2, 65, 66, 7, 15, 2, 2, 66, 68, 7, 16, 2,
	2, 67, 69, 5, 8, 5, 2, 68, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 68,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 7, 17, 2, 2,
	73, 74, 7, 18, 2, 2, 74, 76, 7, 16, 2, 2, 75, 77, 5, 8, 5, 2, 76, 75, 3,
	2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79,
	80, 3, 2, 2, 2, 80, 81, 7, 17, 2, 2, 81, 94, 3, 2, 2, 2, 82, 83, 7, 19,
	2, 2, 83, 84, 5, 6, 4, 2, 84, 85, 7, 20, 2, 2, 85, 87, 7, 16, 2, 2, 86,
	88, 5, 8, 5, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 87, 3, 2, 2,
	2, 89, 90, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 7, 17, 2, 2, 92, 94,
	3, 2, 2, 2, 93, 56, 3, 2, 2, 2, 93, 58, 3, 2, 2, 2, 93, 63, 3, 2, 2, 2,
	93, 82, 3, 2, 2, 2, 94, 9, 3, 2, 2, 2, 11, 13, 18, 32, 34, 54, 70, 78,
	89, 93,
}

var deserializer = antlr.NewATNDeserializer(nil)

var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "'/'", "'+'", "'-'", "'<'", "'>'", "'=='", "'!='", "'skip'",
	"';'", "':='", "'if'", "'then'", "'{'", "'}'", "'else'", "'while'", "'do'",
}

var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "ID", "INT", "BOOL", "WS", "COMMENT",
}

var ruleNames = []string{
	"prog", "aexp", "bexp", "cexp",
}

type whileParser struct {
	*antlr.BaseParser
}

func NewwhileParser(input antlr.TokenStream) *whileParser {
	var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	var sharedContextCache = antlr.NewPredictionContextCache()

	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(whileParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, sharedContextCache)
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "while.g4"

	return this
}

// whileParser tokens.
const (
	whileParserEOF     = antlr.TokenEOF
	whileParserT__0    = 1
	whileParserT__1    = 2
	whileParserT__2    = 3
	whileParserT__3    = 4
	whileParserT__4    = 5
	whileParserT__5    = 6
	whileParserT__6    = 7
	whileParserT__7    = 8
	whileParserT__8    = 9
	whileParserT__9    = 10
	whileParserT__10   = 11
	whileParserT__11   = 12
	whileParserT__12   = 13
	whileParserT__13   = 14
	whileParserT__14   = 15
	whileParserT__15   = 16
	whileParserT__16   = 17
	whileParserT__17   = 18
	whileParserID      = 19
	whileParserINT     = 20
	whileParserBOOL    = 21
	whileParserWS      = 22
	whileParserCOMMENT = 23
)

// whileParser rules.
const (
	whileParserRULE_prog = 0
	whileParserRULE_aexp = 1
	whileParserRULE_bexp = 2
	whileParserRULE_cexp = 3
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = whileParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = whileParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllCexp() []ICexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICexpContext)(nil)).Elem())
	var tst = make([]ICexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICexpContext)
		}
	}

	return tst
}

func (s *ProgContext) Cexp(i int) ICexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICexpContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *whileParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, whileParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<whileParserT__8)|(1<<whileParserT__11)|(1<<whileParserT__16)|(1<<whileParserID))) != 0) {
		{
			p.SetState(8)
			p.Cexp()
		}

		p.SetState(11)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAexpContext is an interface to support dynamic dispatch.
type IAexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAexpContext differentiates from other interfaces.
	IsAexpContext()
}

type AexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAexpContext() *AexpContext {
	var p = new(AexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = whileParserRULE_aexp
	return p
}

func (*AexpContext) IsAexpContext() {}

func NewAexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AexpContext {
	var p = new(AexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = whileParserRULE_aexp

	return p
}

func (s *AexpContext) GetParser() antlr.Parser { return s.parser }

func (s *AexpContext) CopyFrom(ctx *AexpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DivContext struct {
	*AexpContext
}

func NewDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivContext {
	var p = new(DivContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *DivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *DivContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *DivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterDiv(s)
	}
}

func (s *DivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitDiv(s)
	}
}

func (s *DivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type SubContext struct {
	*AexpContext
}

func NewSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubContext {
	var p = new(SubContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *SubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *SubContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *SubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterSub(s)
	}
}

func (s *SubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitSub(s)
	}
}

func (s *SubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulContext struct {
	*AexpContext
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *MulContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitMul(s)
	}
}

func (s *MulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitMul(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarContext struct {
	*AexpContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ID() antlr.TerminalNode {
	return s.GetToken(whileParserID, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumContext struct {
	*AexpContext
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) INT() antlr.TerminalNode {
	return s.GetToken(whileParserINT, 0)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitNum(s)
	}
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

type SumContext struct {
	*AexpContext
}

func NewSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumContext {
	var p = new(SumContext)

	p.AexpContext = NewEmptyAexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AexpContext))

	return p
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *SumContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitSum(s)
	}
}

func (s *SumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitSum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *whileParser) Aexp() (localctx IAexpContext) {
	return p.aexp(0)
}

func (p *whileParser) aexp(_p int) (localctx IAexpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAexpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAexpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, whileParserRULE_aexp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case whileParserINT:
		localctx = NewNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(14)
			p.Match(whileParserINT)
		}

	case whileParserID:
		localctx = NewVarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.Match(whileParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulContext(p, NewAexpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, whileParserRULE_aexp)
				p.SetState(18)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(19)
					p.Match(whileParserT__0)
				}
				{
					p.SetState(20)
					p.aexp(7)
				}

			case 2:
				localctx = NewDivContext(p, NewAexpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, whileParserRULE_aexp)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(22)
					p.Match(whileParserT__1)
				}
				{
					p.SetState(23)
					p.aexp(6)
				}

			case 3:
				localctx = NewSumContext(p, NewAexpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, whileParserRULE_aexp)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(25)
					p.Match(whileParserT__2)
				}
				{
					p.SetState(26)
					p.aexp(5)
				}

			case 4:
				localctx = NewSubContext(p, NewAexpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, whileParserRULE_aexp)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(28)
					p.Match(whileParserT__3)
				}
				{
					p.SetState(29)
					p.aexp(4)
				}

			}

		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IBexpContext is an interface to support dynamic dispatch.
type IBexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBexpContext differentiates from other interfaces.
	IsBexpContext()
}

type BexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBexpContext() *BexpContext {
	var p = new(BexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = whileParserRULE_bexp
	return p
}

func (*BexpContext) IsBexpContext() {}

func NewBexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BexpContext {
	var p = new(BexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = whileParserRULE_bexp

	return p
}

func (s *BexpContext) GetParser() antlr.Parser { return s.parser }

func (s *BexpContext) CopyFrom(ctx *BexpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BoolContext struct {
	*BexpContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.BexpContext = NewEmptyBexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BexpContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(whileParserBOOL, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type LesContext struct {
	*BexpContext
}

func NewLesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LesContext {
	var p = new(LesContext)

	p.BexpContext = NewEmptyBexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BexpContext))

	return p
}

func (s *LesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LesContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *LesContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *LesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterLes(s)
	}
}

func (s *LesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitLes(s)
	}
}

func (s *LesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitLes(s)

	default:
		return t.VisitChildren(s)
	}
}

type NeqContext struct {
	*BexpContext
}

func NewNeqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NeqContext {
	var p = new(NeqContext)

	p.BexpContext = NewEmptyBexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BexpContext))

	return p
}

func (s *NeqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeqContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *NeqContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *NeqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterNeq(s)
	}
}

func (s *NeqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitNeq(s)
	}
}

func (s *NeqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitNeq(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqContext struct {
	*BexpContext
}

func NewEqContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqContext {
	var p = new(EqContext)

	p.BexpContext = NewEmptyBexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BexpContext))

	return p
}

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *EqContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *EqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterEq(s)
	}
}

func (s *EqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitEq(s)
	}
}

func (s *EqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitEq(s)

	default:
		return t.VisitChildren(s)
	}
}

type GtContext struct {
	*BexpContext
}

func NewGtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GtContext {
	var p = new(GtContext)

	p.BexpContext = NewEmptyBexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BexpContext))

	return p
}

func (s *GtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtContext) AllAexp() []IAexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAexpContext)(nil)).Elem())
	var tst = make([]IAexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAexpContext)
		}
	}

	return tst
}

func (s *GtContext) Aexp(i int) IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *GtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterGt(s)
	}
}

func (s *GtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitGt(s)
	}
}

func (s *GtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitGt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *whileParser) Bexp() (localctx IBexpContext) {
	localctx = NewBexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, whileParserRULE_bexp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.aexp(0)
		}
		{
			p.SetState(36)
			p.Match(whileParserT__4)
		}
		{
			p.SetState(37)
			p.aexp(0)
		}

	case 2:
		localctx = NewGtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.aexp(0)
		}
		{
			p.SetState(40)
			p.Match(whileParserT__5)
		}
		{
			p.SetState(41)
			p.aexp(0)
		}

	case 3:
		localctx = NewEqContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.aexp(0)
		}
		{
			p.SetState(44)
			p.Match(whileParserT__6)
		}
		{
			p.SetState(45)
			p.aexp(0)
		}

	case 4:
		localctx = NewNeqContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.aexp(0)
		}
		{
			p.SetState(48)
			p.Match(whileParserT__7)
		}
		{
			p.SetState(49)
			p.aexp(0)
		}

	case 5:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(51)
			p.Match(whileParserBOOL)
		}

	}

	return localctx
}

// ICexpContext is an interface to support dynamic dispatch.
type ICexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCexpContext differentiates from other interfaces.
	IsCexpContext()
}

type CexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCexpContext() *CexpContext {
	var p = new(CexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = whileParserRULE_cexp
	return p
}

func (*CexpContext) IsCexpContext() {}

func NewCexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CexpContext {
	var p = new(CexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = whileParserRULE_cexp

	return p
}

func (s *CexpContext) GetParser() antlr.Parser { return s.parser }

func (s *CexpContext) CopyFrom(ctx *CexpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SkipContext struct {
	*CexpContext
}

func NewSkipContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SkipContext {
	var p = new(SkipContext)

	p.CexpContext = NewEmptyCexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CexpContext))

	return p
}

func (s *SkipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterSkip(s)
	}
}

func (s *SkipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitSkip(s)
	}
}

func (s *SkipContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitSkip(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileContext struct {
	*CexpContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	p.CexpContext = NewEmptyCexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CexpContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) Bexp() IBexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBexpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBexpContext)
}

func (s *WhileContext) AllCexp() []ICexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICexpContext)(nil)).Elem())
	var tst = make([]ICexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICexpContext)
		}
	}

	return tst
}

func (s *WhileContext) Cexp(i int) ICexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICexpContext)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfContext struct {
	*CexpContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.CexpContext = NewEmptyCexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CexpContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) Bexp() IBexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBexpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBexpContext)
}

func (s *IfContext) AllCexp() []ICexpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICexpContext)(nil)).Elem())
	var tst = make([]ICexpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICexpContext)
		}
	}

	return tst
}

func (s *IfContext) Cexp(i int) ICexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICexpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICexpContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitIf(s)
	}
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignContext struct {
	*CexpContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	p.CexpContext = NewEmptyCexpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CexpContext))

	return p
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(whileParserID, 0)
}

func (s *AssignContext) Aexp() IAexpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAexpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAexpContext)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(whileListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case whileVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *whileParser) Cexp() (localctx ICexpContext) {
	localctx = NewCexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, whileParserRULE_cexp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case whileParserT__8:
		localctx = NewSkipContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(whileParserT__8)
		}
		{
			p.SetState(55)
			p.Match(whileParserT__9)
		}

	case whileParserID:
		localctx = NewAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(whileParserID)
		}
		{
			p.SetState(57)
			p.Match(whileParserT__10)
		}
		{
			p.SetState(58)
			p.aexp(0)
		}
		{
			p.SetState(59)
			p.Match(whileParserT__9)
		}

	case whileParserT__11:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Match(whileParserT__11)
		}
		{
			p.SetState(62)
			p.Bexp()
		}
		{
			p.SetState(63)
			p.Match(whileParserT__12)
		}
		{
			p.SetState(64)
			p.Match(whileParserT__13)
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<whileParserT__8)|(1<<whileParserT__11)|(1<<whileParserT__16)|(1<<whileParserID))) != 0) {
			{
				p.SetState(65)
				p.Cexp()
			}

			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(70)
			p.Match(whileParserT__14)
		}
		{
			p.SetState(71)
			p.Match(whileParserT__15)
		}
		{
			p.SetState(72)
			p.Match(whileParserT__13)
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<whileParserT__8)|(1<<whileParserT__11)|(1<<whileParserT__16)|(1<<whileParserID))) != 0) {
			{
				p.SetState(73)
				p.Cexp()
			}

			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(78)
			p.Match(whileParserT__14)
		}

	case whileParserT__16:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)
			p.Match(whileParserT__16)
		}
		{
			p.SetState(81)
			p.Bexp()
		}
		{
			p.SetState(82)
			p.Match(whileParserT__17)
		}
		{
			p.SetState(83)
			p.Match(whileParserT__13)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<whileParserT__8)|(1<<whileParserT__11)|(1<<whileParserT__16)|(1<<whileParserID))) != 0) {
			{
				p.SetState(84)
				p.Cexp()
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(89)
			p.Match(whileParserT__14)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *whileParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *AexpContext = nil
		if localctx != nil {
			t = localctx.(*AexpContext)
		}
		return p.Aexp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *whileParser) Aexp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
