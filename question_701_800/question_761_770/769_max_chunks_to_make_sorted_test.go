package question_761_770

import (
	"testing"
)

func Test_maxChunksToSorted(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{0}, 1},
		{[]int{0, 1, 2, 3}, 4},
		{[]int{1, 2, 0, 3}, 2},
		{[]int{4, 3, 2, 1, 0}, 1},
		{[]int{1, 0, 2, 3, 4}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxChunksToSorted(tt.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
