package question_391_400

import (
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 3},
		{11, 0},
		{10, 1},
		{100000, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findNthDigit(tt.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
