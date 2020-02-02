package question_741_750

import (
	"testing"
)

func Test_dominantIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{0, 1}, 1},
		{[]int{1, 0}, 0},
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := dominantIndex(tt.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
