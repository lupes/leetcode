package question_221_230

import (
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{[][]byte{{'0'}}, 0},
		{[][]byte{{'1'}}, 1},
		{[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 4},
		{[][]byte{{'1', '1', '1', '0', '0'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}, 9},
		{[][]byte{
			{'0', '0', '0', '1'},
			{'1', '1', '0', '1'},
			{'1', '1', '1', '1'},
			{'0', '1', '1', '1'},
			{'0', '1', '1', '1'}}, 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maximalSquare(tt.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
