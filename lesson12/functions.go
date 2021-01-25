package main

import "fmt"

//kelvinToCelsius
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

//celsiusToFahrenheit
func celsiusToFahrenheit(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}

//kelvinToFahrenheit
func kelvinToFahrenheit(k float64) float64 {
	//	k = ((k - 273.15) * 9.0 / 5.0) + 32.0
	//	return k
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	c := kelvinToCelsius(233)
	f := kelvinToFahrenheit(0)

	fmt.Printf("233°K = %.2f°C\n", c)
	fmt.Printf("0°K = %.2f°F\n", f)
}
