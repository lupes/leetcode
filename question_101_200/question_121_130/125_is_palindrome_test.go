package question_121_130

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"test#0", "", true},
		{"test#1", "A man, a plan, a canal: Panama", true},
		{"test#2", "race a car", false},
		{"test#3", "abc", false},
		{"test#4", "aa", true},
		{"test#5", "a1a", true},
		{"test#6", "121", true},
		{"test#7", "0P", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
