package question_1561_1570

import (
	"testing"
)

func Test_getMaxLen(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, -2, -3, 4}, 4},
		{[]int{0, 1, -2, -3, -4}, 3},
		{[]int{-2, -3, -4, 1}, 3},
		{[]int{-1, -2, -3, 0, 1}, 2},
		{[]int{-1, 2}, 1},
		{[]int{1, 2, 3, 5, -6, 4, 0, 10}, 4},
		{[]int{1, 2, 3, 5}, 4},
		{[]int{-32, 13, 26, 7, 19, 0, 32, -27, -34, 26, 39, 22, 23, 23, -11, 0, 14, 10, -14, 1}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getMaxLen(tt.nums); got != tt.want {
				t.Errorf("getMaxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
