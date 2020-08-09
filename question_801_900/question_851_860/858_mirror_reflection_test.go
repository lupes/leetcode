package question_851_860

import (
	"testing"
)

func Test_gcd(t *testing.T) {
	tests := []struct {
		p    int
		q    int
		want int
	}{
		{4, 2, 2},
		{5, 2, 1},
		{28, 21, 7},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := gcd(tt.p, tt.q); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirrorReflection(t *testing.T) {
	tests := []struct {
		p    int
		q    int
		want int
	}{
		{5, 3, 1},
		{2, 1, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mirrorReflection(tt.p, tt.q); got != tt.want {
				t.Errorf("mirrorReflection() = %v, want %v", got, tt.want)
			}
		})
	}
}
