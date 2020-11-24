package question_541_550

import (
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	tests := []struct {
		M    [][]int
		want int
	}{
		{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
		{[][]int{{1, 1, 0}, {1, 1, 1}, {0, 1, 1}}, 1},
		{[][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1},
		{[][]int{
			{1, 0, 0, 1},
			{0, 1, 1, 0},
			{0, 1, 1, 1},
			{1, 0, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findCircleNum(tt.M); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
