package question_1511_1520

import (
	"testing"
)

func Test_numIdenticalPairs(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1, 1, 3}, 4},
		{[]int{1, 1, 1, 1}, 6},
		{[]int{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numIdenticalPairs(tt.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
