package question_1261_1270

import (
	"testing"
)

func Test_countServers(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 0}, {0, 1}}, 0},
		{[][]int{{1, 0}, {1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countServers(tt.grid); got != tt.want {
				t.Errorf("countServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
