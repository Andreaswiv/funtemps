package main

import (
	"flag"
	"fmt"

	"github.com/andreaswiv/funtemps/conv"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var cels float64
var kelv float64
var out string
var funfacts string

func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - fahrenheit, K- Kelvin")

}

func main() {

	flag.Parse()
	if out == "C" && isFlagPassed("F") {
		fmt.Println((fahr), "F er ", conv.FahrenheitToCelsius(fahr), "C")
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Println((cels), "C er ", conv.CelsiusToFahrenheit(cels), "F")
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Println((kelv), "K er ", conv.KelvinToCelsius(kelv), "C")
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Println((cels), "C er ", conv.CelsiusToKelvin(cels), "K")
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Println((kelv), "K er ", conv.KelvinToFahrenheit(kelv), "F")
	}

	if out == "K" && isFlagPassed("F") {
		fmt.Println((fahr), "F er ", conv.FahrenheitToKelvin(fahr), "K")
	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
