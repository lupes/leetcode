package question_51_60

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{" ", 0},
		{"hello", 5},
		{"hello world", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
