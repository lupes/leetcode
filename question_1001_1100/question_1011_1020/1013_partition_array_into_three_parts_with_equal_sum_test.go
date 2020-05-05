package question_1011_1020

import (
	"testing"
)

func Test_canThreePartsEqualSum(t *testing.T) {
	tests := []struct {
		A    []int
		want bool
	}{
		{[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}, true},
		{[]int{0, 2, 1, -6, 6, 7, 9, 1, 2, 0, 1}, false},
		{[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canThreePartsEqualSum(tt.A); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
