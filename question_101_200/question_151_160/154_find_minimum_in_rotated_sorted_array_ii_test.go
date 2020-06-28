package question_151_160

import (
	"testing"
)

func Test_findMin2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 5}, 1},
		{[]int{3, 3, 1}, 1},
		{[]int{1, 3, 3}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMin2(tt.nums); got != tt.want {
				t.Errorf("findMin2() = %v, want %v", got, tt.want)
			}
		})
	}
}
