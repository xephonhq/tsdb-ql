package parser

import (
	"testing"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"reflect"
)

func TestAstBuilder_Visit(t *testing.T) {
	input := antlr.NewInputStream("x := x + 1;")
	lexer := NewwhileLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewwhileParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewPrinterwhileListener(), tree)
	visitor := NewAstBuilder()
	//visitor := BasewhileVisitor{}
	//fmt.Println(visitor)
	//fmt.Println(tree)
	//fmt.Println(visitor.Visit)
	visitor.Visit(tree)
	visitorType := reflect.TypeOf(visitor)
	// FIXME: it has Visit, why it still got a nil error
	for i := 0; i < visitorType.NumMethod(); i++ {
		method := visitorType.Method(i)
		fmt.Println(method.Name)
	}
}
