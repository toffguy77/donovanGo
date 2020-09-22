package tempconv

import "fmt"

// Celsius temperature type
type Celsius float64

// Fahrenheit temperature type
type Fahrenheit float64

// Kelvin temperature type
type Kelvin float64

const (
	// AbsoluteZeroC is absolute zero temperature in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is water freezing temperature in Celsius
	FreezingC Celsius = 0
	// BoilingC is water boiling temperature in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
