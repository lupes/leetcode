package question_1491_1500

import (
	"testing"
)

func Test_numSubseq(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{3, 5, 6, 7}, 9, 4},
		{[]int{3, 3, 6, 8}, 10, 6},
		{[]int{2, 3, 3, 4, 6, 7}, 12, 61},
		{[]int{5, 2, 4, 1, 7, 6, 8}, 16, 127},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numSubseq(tt.nums, tt.target); got != tt.want {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
