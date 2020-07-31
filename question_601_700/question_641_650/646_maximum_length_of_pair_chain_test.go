package question_641_650

import (
	"testing"
)

func Test_findLongestChain(t *testing.T) {
	tests := []struct {
		pairs [][]int
		want  int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 2},
		{[][]int{{1, 2}, {2, 3}, {5, 6}, {3, 4}}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findLongestChain(tt.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
