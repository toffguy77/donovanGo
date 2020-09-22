// Package tempconv converts temperature from different scale into each others
package tempconv

// CToF converts temperature from Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// CToK converts temperature from Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// FToC converts temperature from Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// FToK converts temperature from Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(CToK(FToC(f)))
}

// KToC converts temperature from Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

// KToF converts temperature from Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(CToF(KToC(k)))
}
