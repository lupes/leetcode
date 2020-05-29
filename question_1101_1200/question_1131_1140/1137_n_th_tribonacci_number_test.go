package question_1131_1140

import (
	"testing"
)

func Test_tribonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 4},
		{25, 1389537},
		{40, 12960201916},
		{50, 5742568741225},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := tribonacci(tt.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
