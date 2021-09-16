package question_211_220

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	tests := []struct {
		board []string
		words []string
		want  []string
	}{
		{[]string{"oaan", "etae", "ihkr", "iflv"}, []string{"oath", "pea", "eat", "rain"}, []string{"oath", "eat"}},
		{[]string{"ab", "cd"}, []string{"abcb", "pea", "eat", "rain"}, nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			var board [][]byte
			for _, line := range tt.board {
				board = append(board, []byte(line))
			}
			if got := findWords(board, tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
