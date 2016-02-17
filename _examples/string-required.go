package main

import (
	"fmt"

	"github.com/tj/go-prompt"
)

func main() {
	println("need your name!")
	first := prompt.StringRequired("First name: ")
	last := prompt.StringRequired("Last name: ")
	fmt.Printf("\nHello %s %s\n", first, last)
}
