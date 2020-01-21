package question_721_730

import (
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
