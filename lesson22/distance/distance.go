package main

import (
	"fmt"
	"math"
)

type location struct {
	lat, long float64
}
type world struct {
	radius float64
}

type coordinate struct {
	d, m, s float64
	h       rune
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

//rad函数会将角度转换为弧度
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

//distance使用余弦球面定律计算两个位置之间的距离
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func main() {

	//planets radius
	var mars = world{radius: 3389.5}
	earth := world{radius: 6371.0}

	//Mars locations
	spirit := decimalLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := decimalLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := decimalLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	insight := decimalLocation(coordinate{4, 30, 0, 'N'}, coordinate{135, 54, 0, 'E'})

	sharpMountain := decimalLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMoutain := decimalLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})

	//Earth locations
	london := decimalLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	paris := decimalLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	beijing := decimalLocation(coordinate{39, 55, 0, 'N'}, coordinate{116, 22, 60, 'E'})
	nanning := decimalLocation(coordinate{22, 49, 0, 'N'}, coordinate{108, 18, 60, 'E'})

	fmt.Printf("Spirit to Opportunity %.2f km\n", mars.distance(spirit, opportunity))
	fmt.Printf("Spirit to Curiosity %.2f km\n", mars.distance(spirit, curiosity))
	fmt.Printf("Spirit to InSight %.2f km\n", mars.distance(spirit, insight))
	fmt.Printf("Opportunity to Curiosity %.2f km\n", mars.distance(opportunity, curiosity))
	fmt.Printf("Opportunity to InSight %.2f km\n", mars.distance(opportunity, insight))
	fmt.Printf("Curiosity to InSight %.2f km\n", mars.distance(curiosity, insight))

	//London to Paris
	fmt.Printf("London to Paris %.2f km\n", earth.distance(london, paris))

	//Beijing to Nanning
	fmt.Printf("Beijing to Nanning %.2f km\n", earth.distance(beijing, nanning))

	//Mount Sharp to Olympus
	fmt.Printf("Mount Sharp to Olympus Mons %.2f km\n", mars.distance(sharpMountain, olympusMoutain))
}
