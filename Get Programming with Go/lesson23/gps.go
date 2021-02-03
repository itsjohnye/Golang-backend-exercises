package main

import (
	"fmt"
	"math"
)

type gps struct {
	world
	currentLocation location
	targetLocation  location
}
type name string
type location struct {
	name
	lat, long float64
}
type world struct {
	radius float64
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

func (l location) description() string {
	return fmt.Sprintf("%v (%.2fº, %.2fº)", l.name, l.lat, l.long)
}

func (g gps) distance() float64 {
	return g.world.distance(g.currentLocation, g.targetLocation)
}
func (g gps) message() string {
	return fmt.Sprintf("%v to %v is about %.2f KMs", g.currentLocation.description(), g.targetLocation.description(), g.distance())
}

type rover struct {
	gps
}

func main() {

	mars := world{radius: 3389.5}
	bradbury := location{"Bradbury Landing", -4.5895, 137.4417}
	elysium := location{"Elysium Planitia", 4.5, 135.9}

	gpsInMars := gps{
		world:           mars,
		currentLocation: bradbury,
		targetLocation:  elysium,
	}

	curiosity := rover{
		gps: gpsInMars,
	}

	fmt.Println(curiosity.message())
}
