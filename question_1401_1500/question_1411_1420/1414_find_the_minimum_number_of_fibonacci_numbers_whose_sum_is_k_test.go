package question_1411_1420

import (
	"testing"
)

func Test_findMinFibonacciNumbers(t *testing.T) {
	tests := []struct {
		k    int
		want int
	}{
		{7, 2},
		{2, 1},
		{1, 1},
		{3, 1},
		{10, 2},
		{19, 3},
		{10000000000000, 21},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMinFibonacciNumbers(tt.k); got != tt.want {
				t.Errorf("findMinFibonacciNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
