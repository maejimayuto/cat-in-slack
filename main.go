package main

import (
	"cat-in-slack/cat"
	"fmt"
)

func main() {
	result := cat.GetCat()
	fmt.Printf("%s\n", result)
}
