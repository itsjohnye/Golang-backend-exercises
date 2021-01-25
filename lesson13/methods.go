package main

import "fmt"

type kelvin float64
type celsius float64
type fahrenheit float64

//kelvinToCelsius
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

//kelvinToFahrenheit
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

//celsiusToFahrenheit
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

//celsiusToKelvin
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

//fahrenheitToCelsius
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

//fahrenheitToKelvin
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {
	var k kelvin = 294.0
	var c celsius
	var f fahrenheit

	c = k.celsius()
	f = k.fahrenheit()

	fmt.Printf("%.2f°K = %.2f°C = %.2f°F\n", k, c, f)

	c = 35.0
	k = c.kelvin()
	f = c.fahrenheit()

	fmt.Printf("%.2f°C = %.2f°K = %.2f°F\n", c, k, f)

	f = 89.0
	c = f.celsius()
	k = f.kelvin()

	fmt.Printf("%.2f°F = %.2f°C = %.2f°K\n", f, c, k)
}
