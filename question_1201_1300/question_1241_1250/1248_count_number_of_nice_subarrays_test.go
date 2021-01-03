package question_1241_1250

import (
	"testing"
)

func Test_numberOfSubarrays(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 2, 1, 1}, 3, 2},
		{[]int{2, 4, 6}, 3, 0},
		{[]int{1, 1, 1}, 3, 1},
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 1, 2, 1}, 2, 3},
		{[]int{1, 2, 1, 2, 1}, 2, 4},
		{[]int{1, 2, 1, 2, 1, 2}, 2, 6},
		{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2, 16},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberOfSubarrays(tt.nums, tt.k); got != tt.want {
				t.Errorf("numberOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
