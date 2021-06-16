package question_371_380

import (
	"testing"
)

func Test_wiggleMaxLength(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{0, 0}, 1},
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := wiggleMaxLength(tt.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
