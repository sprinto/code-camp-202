package main

import (
	"fmt"

	"github.com/sprinto/code-camp-2022/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Sinchers")
	fmt.Println(message)
}
