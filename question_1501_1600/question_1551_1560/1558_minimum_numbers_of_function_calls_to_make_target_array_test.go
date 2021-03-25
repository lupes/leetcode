package question_1551_1560

import (
	"testing"
)

func Test_minOperations2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 5}, 5},
		{[]int{2, 2}, 3},
		{[]int{4, 2, 5}, 6},
		{[]int{3, 2, 2, 4}, 7},
		{[]int{2, 4, 8, 16}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minOperations2(tt.nums); got != tt.want {
				t.Errorf("minOperations2() = %v, want %v", got, tt.want)
			}
		})
	}
}
