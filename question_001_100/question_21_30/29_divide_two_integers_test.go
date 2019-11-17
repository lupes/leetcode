package question_0011_0020

import (
	"math"
	"testing"
)

func Test_divide(t *testing.T) {
	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{"test#0", 10, 3, 3},
		{"test#1", 7, -3, -2},
		{"test#2", -7, -3, 2},
		{"test#3", -7, 3, -2},
		{"test#4", math.MinInt32, -1, math.MaxInt32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.dividend, tt.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
