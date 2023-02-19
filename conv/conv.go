package conv

import (
	"fmt"
	"math"
	"strings"
)

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) string {
	result := round((value - 32) * (5.00 / 9.00))
	return addSpaces(result)
}

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) string {
	result := round((value * (9.00 / 5.00)) + 32)
	return addSpaces(result)
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) string {
	result := round(value + 273.15)
	return addSpaces(result)
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) string {
	result := round(value - 273.15)
	return addSpaces(result)
}

// Konverterer Fahrenheit til Kelvin
func FahrenheitToKelvin(value float64) string {
	result := round((value-32)*(5.00/9.00) + 273.15)
	return addSpaces(result)
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) string {
	result := round((value-273.15)*(9.00/5.00) + 32)
	return addSpaces(result)
}

// Runder av til nærmeste heltall og fjerner desimaltall hvis det bare består av nuller
func round(value float64) float64 {
	integerPart := math.Trunc(value)
	decimalPart := value - integerPart

	if decimalPart == 0 {
		return integerPart
	} else {
		return value
	}
}

// Legger til mellomrom på passende steder i tallet og formaterer tall større enn 1000000
func addSpaces(value float64) string {
	strValue := fmt.Sprintf("%.2f", value)

	if len(strValue) <= 3 {
		return strValue
	}

	if value >= 1000000 {
		strValue = fmt.Sprintf("%.0f", math.Trunc(value))
	}

	parts := strings.Split(strValue, ".")
	integerPart := parts[0]
	decimalPart := ""

	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	var result strings.Builder
	for i := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			result.WriteString(" ")
		}
		result.WriteByte(integerPart[i])
	}

	result.WriteString(decimalPart)
	return result.String()
}
