package question_201_210

import (
	"fmt"
	"testing"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	// 1 0001
	// 2 0010
	// 3 0011
	// 4 0100
	// 5 0101
	// 6 0110
	// 7 0111
	// 8 1000
	fmt.Printf("%d\n", 0&1&2&3&4)
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{"test", 5, 7, 4},
		{"test", 4, 8, 0},
		{"test", 63, 64, 0},
		{"test", 30000, 2147483643, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.m, tt.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
