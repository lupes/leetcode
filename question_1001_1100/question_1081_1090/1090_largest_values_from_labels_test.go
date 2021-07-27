package question_1081_1090

import (
	"testing"
)

func Test_largestValsFromLabels(t *testing.T) {
	tests := []struct {
		values     []int
		labels     []int
		num_wanted int
		use_limit  int
		want       int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 1, 2, 2, 3}, 3, 1, 9},
		{[]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2, 12},
		{[]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 1, 16},
		{[]int{9, 8, 8, 7, 6}, []int{0, 0, 0, 1, 1}, 3, 2, 24},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := largestValsFromLabels(tt.values, tt.labels, tt.num_wanted, tt.use_limit); got != tt.want {
				t.Errorf("largestValsFromLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
