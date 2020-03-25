package question_901_910

import (
	"testing"
)

func Test_atMostNGivenDigitSet(t *testing.T) {
	tests := []struct {
		D    []string
		N    int
		want int
	}{
		{[]string{"7"}, 8, 1},
		{[]string{"1", "3", "5", "7"}, 100, 20},
		{[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 899894860, 392738517},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := atMostNGivenDigitSet(tt.D, tt.N); got != tt.want {
				t.Errorf("atMostNGivenDigitSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
