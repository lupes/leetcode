package question_1041_1050

import (
	"testing"
)

func Test_longestDupSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"banana", "ana"},
		{"abcd", ""},
		{"aa", "a"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := longestDupSubstring(tt.s); got != tt.want {
				t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
