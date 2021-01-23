package main

import (
	"fmt"
	"math/rand"
)

func main() {

	const number = 27

	for {
		var n = rand.Intn(100) + 1 //A random integer number is actually a Pseudo random number.
		if n == number {
			fmt.Printf("Well Done! The number is %v. \n", n)
			break //Remenber to break it or it will go infinite.
		} else {
			fmt.Printf("Wrong number %v, guess it agian! \n", n)
		}
	}

}
