package main

import (
	"fmt"

	"io"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/peterh/liner"
	"github.com/xephonhq/tsdb-ql/playground/while/parser"
)

// https://github.com/influxdata/influxdb/blob/master/cmd/influx/cli/cli.go
// test for liner
func loop() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetMultiLineMode(true)
	line.SetCtrlCAborts(true)
	osSignal := make(chan os.Signal, 1)
	for {
		select {
		case <-osSignal:
			// TODO: ctrl + c seems to be handled by liner instead of the loop
			fmt.Println("\nbye~ meow")
			return
		default:
			l, err := line.Prompt("(.w.) ")
			if err == io.EOF {
				fmt.Println("\nbye~ meow")
				return
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Need to parse: %s\n", l)
			// TODO: can't I reuse the lexer? And what do people do when their language have more than one file
			input := antlr.NewInputStream(l)
			lexer := parser.NewwhileLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := parser.NewwhileParser(stream)
			p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true
			tree := p.Prog()
			antlr.ParseTreeWalkerDefault.Walk(parser.NewPrinterwhileListener(), tree)
		}
	}
}

func main() {
	// NOTE: it's ok to use import style like this
	//"./while"
	//w := while.ArithValExp{num: 1}
	//fmt.Println(w)
	loop()
}
