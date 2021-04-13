package question_1601_1610

import (
	"testing"
)

func Test_specialArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 5}, 2},
		{[]int{0, 0}, -1},
		{[]int{0, 4, 3, 0, 4}, 3},
		{[]int{3, 6, 7, 7, 0}, -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := specialArray(tt.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
