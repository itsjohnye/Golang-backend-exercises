package main

import "fmt"

func main() {
	const distance = 236e15
	const lightSpeed = 299792
	const secondsPerYear = 31.536e6
	const lightYears = distance / lightSpeed / secondsPerYear

	fmt.Printf("Canis Major Dwarf Galaxy is %.2f light years away. \n", lightYears)
}
