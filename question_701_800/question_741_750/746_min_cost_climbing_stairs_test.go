package question_741_750

import (
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minCostClimbingStairs(tt.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
