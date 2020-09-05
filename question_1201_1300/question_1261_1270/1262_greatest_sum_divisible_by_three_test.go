package question_1261_1270

import (
	"testing"
)

func Test_maxSumDivThree(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 6, 5, 1, 8}, 18},
		{[]int{4}, 0},
		{[]int{1, 2, 3, 4, 4}, 12},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxSumDivThree(tt.nums); got != tt.want {
				t.Errorf("maxSumDivThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
