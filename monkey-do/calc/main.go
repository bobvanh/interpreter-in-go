package main

import (
	"calc/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This the calc language\n", user.Username)
	fmt.Printf("You can type commands")
	repl.Start(os.Stdin, os.Stdout)
}
