package question_731_740

import (
	"testing"
)

func Test_monotoneIncreasingDigits(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{10, 9},
		{1234, 1234},
		{1231, 1229},
		{1221, 1199},
		{332, 299},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.N); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
