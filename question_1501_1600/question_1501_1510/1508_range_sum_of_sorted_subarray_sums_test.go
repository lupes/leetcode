package question_1501_1510

import (
	"testing"
)

func Test_rangeSum(t *testing.T) {
	tests := []struct {
		nums  []int
		n     int
		left  int
		right int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 4, 1, 5, 13},
		{[]int{1, 2, 3, 4}, 4, 3, 4, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rangeSum(tt.nums, tt.n, tt.left, tt.right); got != tt.want {
				t.Errorf("rangeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
