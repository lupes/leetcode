package question_51_60

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test#0", "", 0},
		{"test#1", " ", 0},
		{"test#2", "hello", 5},
		{"test#3", "hello world", 5},
		{"test#4", "hello world ", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
