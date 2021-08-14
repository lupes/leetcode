package question_1581_1590

import (
	"testing"
)

func Test_unhappyFriends(t *testing.T) {
	tests := []struct {
		n           int
		preferences [][]int
		pairs       [][]int
		want        int
	}{
		{
			4,
			[][]int{{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0}},
			[][]int{{0, 1}, {2, 3}},
			2,
		},
		{
			2,
			[][]int{{1}, {0}},
			[][]int{{0, 1}},
			0,
		},
		{
			4,
			[][]int{{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1}},
			[][]int{{1, 3}, {0, 2}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := unhappyFriends(tt.n, tt.preferences, tt.pairs); got != tt.want {
				t.Errorf("unhappyFriends() = %v, want %v", got, tt.want)
			}
		})
	}
}
