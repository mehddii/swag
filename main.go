package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mehddii/swag/tokenizer"
)

const (
	help string = `
Swag's source code & package manager. 

Usage:
	plug <command> [arguments]

The commands are:
	run		executes Swag source code
	repl	start the interactive repl
	`
	prompt rune = '\U0001F322'
)

func RunFile(path string) {
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}

	Run(string(b))
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%c ", prompt)
		ok := scanner.Scan()
		if !ok {
			return
		}

		Run(scanner.Text())
	}
}

func Run(input string) {
	t := tokenizer.NewTokenizer(input)

	for tok := t.NextToken(); tok.Type != tokenizer.EOF; tok = t.NextToken() {
		fmt.Printf("{\n\ttype: %v,\n\tvalue: %q\n}\n", tok.Type, tok.Literal)
	}
}

func main() {
	args := os.Args

	if len(args) == 3 && args[1] == "run" {
		RunFile(args[2])
	} else if len(args) == 2 && args[1] == "repl" {
		StartRepl()
	} else {
		fmt.Println(help)
		os.Exit(64)
	}
}
