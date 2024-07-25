package main

import (
	"fmt"
	"os"

	"github.com/uniicode-id/uniiscript/usc"
)

func help() {
	fmt.Printf("USC - UniiScript Compiler\n")
	fmt.Printf("\n")
	fmt.Printf("Usage: usc <file.us>\n")
}

func main() {
	args := os.Args
	argsLength := len(args)

	if argsLength != 2 {
		help()
	}

	_ = usc.NewCompiler()
}
