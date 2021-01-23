package main

import "fmt"

func main() {
	var distance = 56000000 //km
	var days = 28

	speed := distance / days / 24 //km per hour

	fmt.Printf("It needs %v km per hour to go to malacandra with 28 days.\n", speed)
}
