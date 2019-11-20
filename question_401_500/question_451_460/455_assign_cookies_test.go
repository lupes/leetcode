package question_451_460

import (
	"testing"
)

func Test_findContentChildren(t *testing.T) {
	tests := []struct {
		g    []int
		s    []int
		want int
	}{
		{[]int{1}, []int{}, 0},
		{[]int{0}, []int{1}, 1},
		{[]int{}, []int{1}, 0},
		{[]int{1}, []int{1}, 1},
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{1, 2, 2, 3, 4, 3}, []int{1, 2, 3}, 3},
		{[]int{1, 2, 2, 3, 4, 3, 2}, []int{1, 2, 3, 3}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findContentChildren(tt.g, tt.s); got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
