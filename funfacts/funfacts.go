package funfacts

import (
	"flag"
	"fmt"
)

var fahrSun string
var celsSun string
var kelvSun string
var t string
var funfacts string

func init() {

	flag.StringVar(&fahrSun, "F", "9944.33°F", "temperatur i grader fahrenheit")
	flag.StringVar(&kelvSun, "F", "5780°K", "temperatur i grader kelvin")
	flag.StringVar(&celsSun, "C", "5506.85°C.", "temperatur i grader celsius")

	flag.StringVar(&t, "t", "C", "beregne temperatur i C - celsius, F - fahrenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")

}

func main() {

	flag.Parse()
	if t == "C" && isFlagPassed("sun") {
		fmt.Println("Temperatur på ytre lag av Solen", (celsSun))
	}

	if t == "F" && isFlagPassed("sun") {
		fmt.Println("Temperatur på ytre lag av Solen", (fahrSun))
	}

	if t == "K" && isFlagPassed("sun") {
		fmt.Println("Temperatur på ytre lag av Solen", (kelvSun))
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
