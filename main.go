package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/grobinson-grafana/matchers"
)

var (
	justLex bool
)

func parseFlags() {
	flag.BoolVar(&justLex, "just-lex", false, "skip parsing and print the tokens from the lexer")
	flag.Parse()
}

func printMatchers(s string) {
	m, err := matchers.Parse(strings.TrimSpace(s))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintln(os.Stdout, m)
	}
}

func printTokens(s string) {
	l := matchers.NewLexer(strings.TrimSpace(s))
	for {
		tok, err := l.Scan()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		} else if tok != (matchers.Token{}) {
			fmt.Fprintln(os.Stdout, tok)
		} else {
			break
		}
	}
}

func main() {
	parseFlags()

	r := bufio.NewReader(os.Stdin)
	for {
		in, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			fmt.Fprintf(os.Stderr, "unexpected error: %s\n", err)
		}
		if justLex {
			printTokens(in)
		} else {
			printMatchers(in)
		}
	}
}
