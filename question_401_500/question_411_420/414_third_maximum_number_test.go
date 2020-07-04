package question_411_420

import (
	"testing"
)

func Test_thirdMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{2, 2, 1}, 2},
		{[]int{2, 1}, 2},
		{[]int{2, 2, 3, 1}, 1},
		{[]int{1, 2, 2, 5, 3, 5}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := thirdMax(tt.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
