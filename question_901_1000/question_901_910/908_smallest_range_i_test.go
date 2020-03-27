package question_901_910

import (
	"testing"
)

func Test_smallestRangeI(t *testing.T) {
	tests := []struct {
		A    []int
		K    int
		want int
	}{
		{[]int{1}, 0, 0},
		{[]int{0, 10}, 2, 6},
		{[]int{1, 3, 6}, 3, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := smallestRangeI(tt.A, tt.K); got != tt.want {
				t.Errorf("smallestRangeI() = %v, want %v", got, tt.want)
			}
		})
	}
}
