package question_1531_1540

import (
	"testing"
)

func Test_getWinner(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{2, 1, 3, 5, 4, 6, 7}, 2, 5},
		{[]int{3, 2, 1}, 10, 3},
		{[]int{1, 9, 8, 2, 3, 7, 6, 4, 5}, 7, 9},
		{[]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 10000000000, 99},
		{[]int{1, 25, 35, 42, 68, 70}, 1, 25},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getWinner(tt.arr, tt.k); got != tt.want {
				t.Errorf("getWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
