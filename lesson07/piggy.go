package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var piggyBank = 0

	for piggyBank < 2000 {

		switch rand.Intn(3) {
		case 0:
			piggyBank += 5
		case 1:
			piggyBank += 10
		case 2:
			piggyBank += 25
		}

		dollars := piggyBank / 100
		cents := piggyBank % 100
		fmt.Printf("$%d.%02d \n", dollars, cents) //"%02d"表示: '%'格式标识符；'0'填充；'2'位数；'d'十进制格式。

	}
}
