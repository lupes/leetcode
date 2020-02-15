package question_831_840

import (
	"testing"
)

func Test_findReplaceString(t *testing.T) {
	tests := []struct {
		S       string
		indexes []int
		sources []string
		targets []string
		want    string
	}{
		{"abcd", []int{1}, []string{"bc"}, []string{"eee"}, "aeeed"},
		{"abcd", []int{1}, []string{"bcde"}, []string{"eee"}, "abcd"},
		{"abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}, "eeebffff"},
		{"abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}, "eeecd"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findReplaceString(tt.S, tt.indexes, tt.sources, tt.targets); got != tt.want {
				t.Errorf("findReplaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
