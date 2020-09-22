package tempconv

import (
	"testing"
)

func TestCelsius_String(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want string
	}{
		{"Boiling point of water", 100, "100°C"},
		{"Melting point of ice", 0, "0°C"},
		{"Typical human body temperature", 37.0, "37°C"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Celsius.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFahrenheit_String(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want string
	}{
		{"Boiling point of water", 212, "212°F"},
		{"Melting point of ice", 32, "32°F"},
		{"Typical human body temperature", 98.6, "98.6°F"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("Fahrenheit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_String(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want string
	}{
		{"Boiling point of water", 373.15, "373.15K"},
		{"Melting point of ice", 273.15, "273.15K"},
		{"Typical human body temperature", 310.15, "310.15K"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Kelvin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
