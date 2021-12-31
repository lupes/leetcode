package question_841_850

import (
	"testing"
)

func Test_isNStraightHand(t *testing.T) {
	tests := []struct {
		hand []int
		W    int
		want bool
	}{
		{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 4, false},
		{[]int{1, 2, 3}, 1, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isNStraightHand(tt.hand, tt.W); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
