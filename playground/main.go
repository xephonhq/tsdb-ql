package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func main() {
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
