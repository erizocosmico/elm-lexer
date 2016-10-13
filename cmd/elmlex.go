package main

import (
	"flag"
	"fmt"
	"os"

	elmlexer "github.com/mvader/elm-lexer"
)

var help = flag.Bool("help", false, "display help")

func main() {
	flag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	lexer := elmlexer.New(os.Stdin)
	go lexer.Run()

	for {
		t, ok := lexer.Next()
		if !ok {
			break
		}

		fmt.Printf(
			"LINE: %4d POS: %4d TYPE: %-30s %s\n",
			t.Line,
			t.LinePos,
			t.Type,
			t.Value,
		)
	}
}

const helpText = `Display a list of tokens with their properties

usage: elmlex < /path/to/file.elm`

func printUsage() {
	fmt.Println(helpText)
}
