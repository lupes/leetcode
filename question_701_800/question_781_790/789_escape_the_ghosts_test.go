package question_781_790

import (
	"testing"
)

func Test_escapeGhosts(t *testing.T) {
	tests := []struct {
		ghosts [][]int
		target []int
		want   bool
	}{
		{[][]int{{1, 0}, {0, 3}}, []int{0, 1}, true},
		{[][]int{{1, 0}}, []int{2, 0}, false},
		{[][]int{{2, 0}}, []int{1, 0}, false},
		{[][]int{{5, 0}, {-10, -2}, {0, -5}, {-2, -2}, {-7, 1}}, []int{7, 7}, false},
		{[][]int{{-1, 0}, {0, 1}, {-1, 0}, {0, 1}, {-1, 0}}, []int{0, 0}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := escapeGhosts(tt.ghosts, tt.target); got != tt.want {
				t.Errorf("escapeGhosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
