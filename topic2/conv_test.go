// Package tempconv converts temperature from different scale into each others
package tempconv

import (
	"testing"
)

func TestCToF(t *testing.T) {
	tests := []struct {
		name string
		args Celsius
		want Fahrenheit
	}{
		{"Boiling point of water", 100, 212},
		{"Melting point of ice", 0, 32},
		{"Typical human body temperature", 37.0, 98.6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToF(tt.args); got != tt.want {
				t.Errorf("CToF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCToK(t *testing.T) {
	tests := []struct {
		name string
		args Celsius
		want Kelvin
	}{
		{"Boiling point of water", 100, 373.15},
		{"Melting point of ice", 0, 273.15},
		{"Typical human body temperature", 37.0, 310.15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToK(tt.args); got != tt.want {
				t.Errorf("CToK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		name string
		args Fahrenheit
		want Celsius
	}{
		{"Boiling point of water", 212, 100},
		{"Melting point of ice", 32, 0},
		{"Typical human body temperature", 98.6, 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FToC(tt.args); got != tt.want {
				t.Errorf("FToC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFToK(t *testing.T) {
	tests := []struct {
		name string
		args Fahrenheit
		want Kelvin
	}{
		{"Boiling point of water", 212, 373.15},
		{"Melting point of ice", 32, 273.15},
		{"Typical human body temperature", 98.6, 310.15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FToK(tt.args); got != tt.want {
				t.Errorf("FToK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKToC(t *testing.T) {
	tests := []struct {
		name string
		args Kelvin
		want Celsius
	}{
		{"Boiling point of water", 373.15, 100},
		{"Melting point of ice", 273.15, 0},
		{"Typical human body temperature", 310.15, 37},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToC(tt.args); got != tt.want {
				t.Errorf("KToC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKToF(t *testing.T) {
	tests := []struct {
		name string
		args Kelvin
		want Fahrenheit
	}{
		{"Boiling point of water", 373.15, 212},
		{"Melting point of ice", 273.15, 32},
		{"Typical human body temperature", 310.15, 98.6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KToF(tt.args); got != tt.want {
				t.Errorf("KToF() = %v, want %v", got, tt.want)
			}
		})
	}
}
