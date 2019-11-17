package question_341_350

import "testing"

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test", "hello", "holle"},
		{"test", "leetcode", "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
