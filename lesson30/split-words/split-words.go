package main

import (
	"fmt"
	"strings"
)

func sourceInput(downstream chan string) {
	for _, v := range []string{"hello yes no world", "a bad apple", "goodbye all hello"} {
		downstream <- v
	}
	close(downstream)
}

func splitHandling(upstream, downstream chan string) {
	for i := range upstream {
		for _, word := range strings.Fields(i) {
			downstream <- word
		}
	}
	close(downstream)
}

func printString(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceInput(c0)
	go splitHandling(c0, c1)
	printString(c1)
}
