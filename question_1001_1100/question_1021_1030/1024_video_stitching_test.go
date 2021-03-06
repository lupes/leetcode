package question_1021_1030

import (
	"testing"
)

func Test_videoStitching(t *testing.T) {
	tests := []struct {
		clips [][]int
		T     int
		want  int
	}{
		{[][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10, 3},
		{[][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9, 3},
		{[][]int{{0, 1}, {1, 2}}, 5, -1},
		{[][]int{{0, 4}, {2, 8}}, 5, 2},
		{[][]int{{0, 4}, {2, 5}}, 6, -1},
		{[][]int{{0, 2}, {4, 8}}, 5, -1},
		{[][]int{{5, 7}, {1, 8}, {0, 0}, {2, 3}, {4, 5}, {0, 6}, {5, 10}, {7, 10}}, 5, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := videoStitching(tt.clips, tt.T); got != tt.want {
				t.Errorf("videoStitching() = %v, want %v", got, tt.want)
			}
		})
	}
}
