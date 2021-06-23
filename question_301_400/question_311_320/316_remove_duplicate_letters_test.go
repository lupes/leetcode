package question_311_320

import (
	"testing"
)

func Test_removeDuplicateLetters(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"bcabc", "abc"},
		{"abacb", "abc"},
		{"cbacdcbc", "acdb"},
		{"alngkakgalkanafoqklga", "agklnfoq"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeDuplicateLetters(tt.s); got != tt.want {
				t.Errorf("removeDuplicateLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
