package main

import (
	"fmt"
	"orca/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("**Wistles** %s! This is the Orca Programming language!\n",
		user.Username)
	fmt.Printf("Pass us some code.\n")
	repl.Start(os.Stdin, os.Stdout)
}
