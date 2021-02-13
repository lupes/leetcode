package question_1451_1460

import (
	"testing"
)

func Test_isPrefixOfWord(t *testing.T) {
	tests := []struct {
		sentence   string
		searchWord string
		want       int
	}{
		{"i love eating burger", "burg", 4},
		{"this problem is an easy problem", "pro", 2},
		{"i am tired", "you", -1},
		{"i use triple pillow", "pill", 4},
		{"hello from the other side", "they", -1},
		{"love errichto jonathan dumb", "dumb", 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isPrefixOfWord(tt.sentence, tt.searchWord); got != tt.want {
				t.Errorf("isPrefixOfWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
