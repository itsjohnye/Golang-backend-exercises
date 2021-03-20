package main

import "fmt"

func main() {

	sayhi("Alice", 10)

}

func sayhi(name string, time int) {

	time -= 1

	fmt.Printf("Hi, %v\n", name)

	if time > 0 {
		sayhi(name, time)
	} else {
		return
	}

}
