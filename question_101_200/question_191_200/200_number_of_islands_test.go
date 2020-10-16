package question_191_200

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{[][]byte{{'1', '1', '1', '1', '0'}}, 1},
		{[][]byte{{'1', '1', '0', '1', '0'}}, 2},
		{[][]byte{{'1', '1', '0', '1', '0'}, {'1', '1', '1', '1', '0'}}, 1},
		{[][]byte{{'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '1', '0', '1'}}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numIslands(tt.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
