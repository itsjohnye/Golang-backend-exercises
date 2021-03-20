package main

import (
	"fmt"
	"math"
)

func main() {

	//fmt.Println(hypot(3, 4))
	fmt.Println(hypot2(3, 4))
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func hypot2(x, y float64) (r float64) {
	r = math.Sqrt(x*x + y*y)
	return
}
