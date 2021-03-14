package main

import "fmt"

func main() {
	withoutExcalmatoryMark()
	withExcalmatoryMark()
}

func withoutExcalmatoryMark() {
	fmt.Println("without Excalmatory Mark: ")
	x := "hello!!you"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c\n", x)
		}
	}
}

func withExcalmatoryMark() {
	fmt.Println("with Excalmatory Mark: ")
	x := "hello!!you"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c\n", x)
	}
}
