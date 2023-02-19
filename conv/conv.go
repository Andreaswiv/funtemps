package conv

import (
	"math"
)

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	result := round((value - 32) * (5.00 / 9.00))
	return result
}

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	result := round((value * (9.00 / 5.00)) + 32)
	return result
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {
	result := round(value + 273.15)
	return result
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {
	result := round(value - 273.15)
	return result
}

// Konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(value float64) float64 {
	result := round((value-32)*(5.00/9.00) + 273.15)
	return result
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	result := round((value-273.15)*(9.00/5.00) + 32)
	return result
}

// Runder av til nærmeste heltall og fjerner desimaltall hvis det bare består av nuller
func round(value float64) float64 {
	return math.Round(value*100) / 100
}
