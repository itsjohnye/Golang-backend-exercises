package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:] //ternimal arguments
	var yesNo bool

	if len(args) == 0 {
		fmt.Println("Please type with argument 'true/yes/1' or 'false/no/0'.")
	}

	for _, c := range args {
		switch c {
		case "true", "yes", "1":
			yesNo = true
		case "false", "no", "0":
			yesNo = false
		default:
			fmt.Println("Invalid arguments", c)
			return
		}
		fmt.Printf("The out put is %v \n", yesNo)
	}

}
