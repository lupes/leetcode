package question_1291_1300

import (
	"testing"
)

func Test_isPossibleDivide(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		//{[]int{1, 2, 3, 3, 4, 4, 5, 6}, 4, true},
		//{[]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3, true},
		//{[]int{3, 3, 2, 2, 1, 1}, 3, true},
		{[]int{3, 3, 2, 2, 1, 1}, 2, true},
		//{[]int{1, 2, 3, 4}, 3, false},
		//{[]int{1, 2, 3, 5}, 4, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isPossibleDivide(tt.nums, tt.k); got != tt.want {
				t.Errorf("isPossibleDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}
