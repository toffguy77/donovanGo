package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/toffguy77/donovanGo/topic2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "converter: %v\n", err)
			os.Exit(1)
		}
		convertInput(t)
	}
}

func convertInput(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)
	fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
}
