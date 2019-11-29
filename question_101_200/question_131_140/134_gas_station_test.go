package question_131_140

import (
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas  []int
		cost []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canCompleteCircuit(tt.gas, tt.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
