package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Selyss/AutoCal/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is AutoCal, your smart calendar automation tool.\n", // TODO: add cobra
		user.Username)
	fmt.Printf("Start by typing in commands:\n")
	repl.Start(os.Stdin, os.Stdout)
}
