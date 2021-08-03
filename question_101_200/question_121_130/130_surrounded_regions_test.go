package question_121_130

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	tests := []struct {
		board [][]byte
		want  [][]byte
	}{
		{
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			[][]byte{{'X'}},
			[][]byte{{'X'}},
		},
		{
			[][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}},
			[][]byte{{'X', 'O', 'X'}, {'X', 'O', 'X'}, {'X', 'O', 'X'}},
		},
		{
			[][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}},
			[][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			solve(tt.board)
			if !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.board, tt.want)
			}
		})
	}
}
