package question_421_430

import (
	"testing"
)

func Test_findMaximumXOR(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 10, 5, 25, 2, 8}, 28},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findMaximumXOR(tt.nums); got != tt.want {
				t.Errorf("findMaximumXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
