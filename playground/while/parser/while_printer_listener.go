package parser

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// PrinterwhileListener just print
type PrinterwhileListener struct {
	// TODO: why embedding like this does not have problem
	// Because BasewhileListener does not embed other things
	*BasewhileListener
	a int
}

// NewPrinterwhileListener returns a new PrinterwhileListener
func NewPrinterwhileListener() *PrinterwhileListener {
	return new(PrinterwhileListener)
}

// EnterEveryRule simply print things out
func (l *PrinterwhileListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
