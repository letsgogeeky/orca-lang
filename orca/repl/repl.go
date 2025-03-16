package repl

import (
	"bufio"
	"fmt"
	"io"
	"orca/lexer"
	"orca/token"
	"os"
)

// REPL management commands
const (
	LEAVE = "leave()"
	HELP  = "help()"
)

func help(out io.Writer, commands []map[string]string) {
	fmt.Fprintln(out, "Available commands:")
	for _, cmd := range commands {
		for k, v := range cmd {
			fmt.Fprintf(out, "  %s: %s\n", k, v)
		}
	}
}

func leave() {
	os.Exit(0)
}

var commands = []map[string]string{
	{LEAVE: "leave(): Orca leaves the Pod."},
	{HELP: "help(): show this help"},
}

// REPL commands

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == LEAVE {
			leave()
			continue
		}
		if line == HELP {
			help(out, commands)
			continue
		}

		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}
