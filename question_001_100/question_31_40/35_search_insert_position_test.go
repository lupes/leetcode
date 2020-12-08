package question_31_40

import "testing"

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{}, 1, 0},
		{[]int{1, 3, 5, 7}, 0, 0},
		{[]int{1, 3, 5, 6}, 1, 0},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 3, 1},
		{[]int{1, 3, 5, 6}, 4, 2},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 6, 3},
		{[]int{1, 3, 5, 6}, 7, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
