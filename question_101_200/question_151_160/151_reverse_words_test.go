package question_151_160

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test", "the sky is blue", "blue is sky the"},
		{"test", "  hello world!  ", "world! hello"},
		{"test", "a good   example", "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
