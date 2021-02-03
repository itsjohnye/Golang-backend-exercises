package main

import "fmt"

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for i, c := range alpha {
		fmt.Printf("%v-%c \n", i, c)
	}
}
