package question_1481_1490

import (
	"testing"
)

func Test_findLeastNumOfUniqueInts(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{5, 5, 4}, 1, 1},
		{[]int{4, 3, 1, 1, 3, 3, 2}, 3, 2},
		{[]int{2, 4, 1, 8, 3, 5, 1, 3}, 3, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findLeastNumOfUniqueInts(tt.arr, tt.k); got != tt.want {
				t.Errorf("findLeastNumOfUniqueInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
