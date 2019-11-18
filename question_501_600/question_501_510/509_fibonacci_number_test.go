package question_501_510

import (
	"testing"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := fib(tt.N); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
