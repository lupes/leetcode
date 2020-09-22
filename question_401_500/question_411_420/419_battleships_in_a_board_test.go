package question_411_420

import (
	"testing"
)

func Test_countBattleships(t *testing.T) {
	tests := []struct {
		board [][]byte
		want  int
	}{
		{[][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countBattleships(tt.board); got != tt.want {
				t.Errorf("countBattleships() = %v, want %v", got, tt.want)
			}
		})
	}
}
