package question_791_800

import (
	"testing"
)

// OXX
// XOX
// OXO

func Test_validTicTacToe(t *testing.T) {
	tests := []struct {
		board []string
		want  bool
	}{
		{[]string{"O  ", "   ", "   "}, false},
		{[]string{"XOX", " X ", "   "}, false},
		{[]string{"XXX", "   ", "OOO"}, false},
		{[]string{"XOX", "O O", "XOX"}, true},
		{[]string{"XOX", "OXO", "XOX"}, true},
		{[]string{"XXX", "XOO", "OO "}, false},
		{[]string{"OXX", "XOX", "OXO"}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := validTicTacToe(tt.board); got != tt.want {
				t.Errorf("validTicTacToe() = %v, want %v", got, tt.want)
			}
		})
	}
}
