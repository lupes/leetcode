package question_661_670

import (
	"testing"
)

func Test_maximumSwap(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{0, 0},
		{1, 1},
		{12, 21},
		{2736, 7236},
		{9973, 9973},
		{19973, 99173},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maximumSwap(tt.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
