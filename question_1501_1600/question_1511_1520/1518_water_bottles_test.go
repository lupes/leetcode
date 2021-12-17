package question_1511_1520

import (
	"testing"
)

func Test_numWaterBottles(t *testing.T) {
	tests := []struct {
		numBottles  int
		numExchange int
		want        int
	}{
		{9, 3, 13},
		{15, 4, 19},
		{5, 5, 6},
		{2, 3, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numWaterBottles(tt.numBottles, tt.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
