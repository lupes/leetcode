package question_1691_1700

import (
	"testing"
)

func Test_maximumUniqueSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 2, 4, 5, 6}, 17},
		{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.nums); got != tt.want {
				t.Errorf("maximumUniqueSuxbarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
