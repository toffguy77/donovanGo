package popcount

import "testing"

func TestPopCount(t *testing.T) {
	tests := []struct {
		name string
		args uint
		want int
	}{
		{"1", 1, 0},
		{"2", 2, 1},
		{"1234", 1234, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount(tt.args); got != tt.want {
				t.Errorf("PopCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
