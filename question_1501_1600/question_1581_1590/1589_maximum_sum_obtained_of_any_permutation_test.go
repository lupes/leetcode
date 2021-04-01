package question_1581_1590

import (
	"testing"
)

func Test_maxSumRangeQuery(t *testing.T) {
	tests := []struct {
		nums     []int
		requests [][]int
		want     int
	}{
		{[]int{1, 2, 3, 4, 5}, [][]int{{1, 3}, {0, 1}}, 19},
		{[]int{1, 2, 3, 4, 5, 6}, [][]int{{0, 1}}, 11},
		{[]int{1, 2, 3, 4, 5, 10}, [][]int{{0, 2}, {1, 3}, {1, 1}}, 47},
		{[]int{1000000000000000, 1000000000000000}, [][]int{{0, 1}}, 47},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxSumRangeQuery(tt.nums, tt.requests); got != tt.want {
				t.Errorf("maxSumRangeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
