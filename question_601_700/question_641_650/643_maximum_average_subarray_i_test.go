package question_641_650

import (
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 2, 3, 4}, 4, 2.5},
		{[]int{1, 2, 3, 4}, 3, 3},
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMaxAverage(tt.nums, tt.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
