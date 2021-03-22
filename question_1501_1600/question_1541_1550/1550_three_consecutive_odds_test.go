package question_1541_1550

import (
	"testing"
)

func Test_threeConsecutiveOdds(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{1, 2, 3}, false},
		{[]int{1, 2, 3, 5, 7}, true},
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 12}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := threeConsecutiveOdds(tt.arr); got != tt.want {
				t.Errorf("threeConsecutiveOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
