package question_1981_1990

import (
	"testing"
)

func Test_minimumDifference(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{90}, 1, 0},
		{[]int{9, 4, 1, 7}, 2, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minimumDifference(tt.nums, tt.k); got != tt.want {
				t.Errorf("minimumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
