package question_1001_1010

import (
	"testing"
)

func Test_largestSumAfterKNegations(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want int
	}{
		{[]int{4, 2, 3}, 1, 5},
		{[]int{4, 2, 3}, 2, 9},
		{[]int{3, -1, 0, 2}, 3, 6},
		{[]int{3, -1, 0, 2}, 4, 6},
		{[]int{2, -3, -1, 5, -4}, 2, 13},
		{[]int{-2, 9, 9, 8, 4}, 5, 32},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := largestSumAfterKNegations(tt.A, tt.K); got != tt.want {
				t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
			}
		})
	}
}
