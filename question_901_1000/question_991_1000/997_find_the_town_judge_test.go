package question_991_1000

import (
	"testing"
)

func Test_findJudge(t *testing.T) {
	tests := []struct {
		N     int
		trust [][]int
		want  int
	}{
		{2, [][]int{{1, 2}}, 2},
		{3, [][]int{{1, 3}, {2, 3}}, 3},
		{3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
		{3, [][]int{{1, 2}, {2, 3}}, -1},
		{4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findJudge(tt.N, tt.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
