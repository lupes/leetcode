package question_411_420

import (
	"testing"
)

func Test_numberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{1, 2, 3, 5, 7, 8, 9, 10, 11}, 8},
		{[]int{1, 2, 3, 5, 7, 8, 9, 10, 11, 13}, 8},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberOfArithmeticSlices(tt.A); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
