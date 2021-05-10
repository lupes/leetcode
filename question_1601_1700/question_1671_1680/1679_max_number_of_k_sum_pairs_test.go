package question_1671_1680

import (
	"testing"
)

func Test_maxOperations(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		//{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxOperations(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
