package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func decimalLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1

	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {

	spirit := decimalLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := decimalLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := decimalLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	insight := decimalLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	fmt.Printf("Spirit: %.4f\n", spirit)
	fmt.Printf("Opportunity: %.4f\n", opportunity)
	fmt.Printf("Curiosity: %.4f\n", curiosity)
	fmt.Printf("InSight: %.4f\n", insight)

}
