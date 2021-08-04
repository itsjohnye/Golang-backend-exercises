package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//couter
	go func() {
		defer close(naturals)
		for i := 0; i < 10; i++ {
			naturals <- i
		}
	}()

	//squarer
	go func() {
		defer close(squares)
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
	}()

	//printer
	for p := range squares {
		fmt.Println(p)
	}
}

//i -> naturals -> x -> squares ->p
