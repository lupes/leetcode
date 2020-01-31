package question_731_740

import (
	"testing"
)

func Test_deleteAndEarn(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
		{[]int{2, 2, 3, 3, 3, 4, 4}, 12},
		{[]int{1, 1, 1, 2, 4, 5, 5, 5, 6}, 18},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := deleteAndEarn(tt.nums); got != tt.want {
				t.Errorf("deleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
