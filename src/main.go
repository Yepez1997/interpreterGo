package main

import "os/user"

import (
	"fmt"
	repl "github.com/Yepez1997/interpreterGo/src/repl"
	"os"
)

// start the repl process
func main() {
	_, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello !!!!")
	fmt.Printf("Feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
