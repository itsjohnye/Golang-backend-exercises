package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Map(add1, "123"))
}

func add1(r rune) rune {
	return r + 1
}
