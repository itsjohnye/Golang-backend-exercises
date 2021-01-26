package main

import "fmt"

func main() {
	slice := []string{}
	lastCapacity := cap(slice) //The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

	for i := 0; i < 10000; i++ {
		slice = append(slice, "An element")
		if cap(slice) != lastCapacity {
			fmt.Printf("cap = %v, len = %v\n", cap(slice), len(slice))
			lastCapacity = cap(slice)
		}
	}
}
