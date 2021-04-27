package question_1651_1660

import (
	"testing"
)

func Test_closeStrings(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"cabbba", "aabbss", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := closeStrings(tt.word1, tt.word2); got != tt.want {
				t.Errorf("closeStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
