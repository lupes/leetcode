package question_1491_1500

import (
	"testing"
)

func Test_canArrange(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5, true},
		{[]int{1, 2, 3, 4, 5, 6}, 7, true},
		{[]int{1, 2, 3, 4, 5, 6}, 10, false},
		{[]int{-10, 10}, 10, true},
		{[]int{-1, 1, -2, 2, -3, 3, -4, 4}, 3, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canArrange(tt.arr, tt.k); got != tt.want {
				t.Errorf("canArrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
