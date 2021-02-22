package question_1461_1470

import (
	"testing"
)

func Test_maxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 2}, 12},
		{[]int{1, 5, 4, 5}, 16},
		{[]int{3, 7}, 12},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
