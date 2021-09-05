package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)



func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey PL!\n", user.Username)
	fmt.Printf("Type anything in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}