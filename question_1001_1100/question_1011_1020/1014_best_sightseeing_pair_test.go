package question_1011_1020

import (
	"testing"
)

func Test_maxScoreSightseeingPair(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{8, 1, 5, 2, 6}, 11},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.A); got != tt.want {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
