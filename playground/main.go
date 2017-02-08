package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func main() {
	// NOTE: it's ok to use import style like this
	//"./while"
	//w := while.ArithValExp{num: 1}
	//fmt.Println(w)
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	s, err := line.Prompt("$ ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
