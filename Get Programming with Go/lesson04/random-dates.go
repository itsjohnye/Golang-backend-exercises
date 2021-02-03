package main

import (
	"fmt"
	"math/rand"
)

var ear = "AD"

func main() {
	for i := 0; i < 10; i++ {
		year := rand.Intn(5000) + 1
		leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2: //case condition with colon
			daysInMonth = 28 //just need an 'if' condition in this case
			if leap {
				daysInMonth = 29
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Println(ear, year, month, day)
	}
}
