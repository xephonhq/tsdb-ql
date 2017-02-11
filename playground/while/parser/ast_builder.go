package parser

import (
	"github.com/xephonhq/tsdb-ql/playground/while/ast"
	"fmt"
)

type AstBuilder struct {
	*BasewhileVisitor
	root ast.CommandSeqExp
}

func NewAstBuilder() *AstBuilder {
	// TODO: do I need to initialize root exp explicitly?
	return new(AstBuilder)
}

func (b *AstBuilder) VisitProg(ctx *ProgContext) interface{} {
	fmt.Println("vist prog!")
	return b.VisitChildren(ctx)
}

//func (b *AstBuilder) Visit(tree antlr.ParseTree) interface{} {
//	return nil
//}