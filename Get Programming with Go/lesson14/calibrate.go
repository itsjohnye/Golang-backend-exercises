package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {

	var tempOffset kelvin = 10

	sensor := calibrate(fakeSensor, tempOffset)

	for count := 0; count < 10; count++ {
		fmt.Println(sensor())
	}

}
