package question_71_80

import "testing"

func Test_exist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{"test", nil, "abc", false},
		{"test", [][]byte{{'a'}}, "a", true},
		{"test", [][]byte{{'a', 'a'}}, "aaa", false},
		{"test", [][]byte{{'a', 'b'}}, "ab", true},
		{"test", [][]byte{{'a', 'b'}}, "abc", false},
		{"test", [][]byte{{'a'}, {'b'}}, "ab", true},
		{"test", [][]byte{{'a'}, {'a'}, {'b'}}, "ab", true},
		{"test", [][]byte{{'a'}, {'a'}, {'c'}, {'b'}}, "ab", false},
		{"test", [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}, "eb", true},
		{"test", [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}, "ef", true},
		{"test", [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}, "ed", true},
		{"test", [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}}, "eh", true},
		{"test", [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{"test", [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{"test", [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
		{"test", [][]byte{{'a', 'b', 'c'}, {'a', 'b', 'a'}}, "ac", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.board, tt.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
