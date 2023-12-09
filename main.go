package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is AutoCal, your smart calendar automation tool.\n", // TODO: add cobra
		user.Username)
}
