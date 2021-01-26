package main

import "fmt"

type planets []string

func (p planets) terraform() planets {
	for i := range p {
		p[i] = "New " + p[i]
	}
	return p
}

func main() {
	var pl planets = []string{
		"Mars", "Uranus", "Neptune",
	}
	fmt.Println(pl)
	fmt.Println(pl.terraform())

}

//original
/*
type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)
}
*/
