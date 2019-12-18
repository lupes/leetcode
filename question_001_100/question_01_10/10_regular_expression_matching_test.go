package question_01_10

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"", "", true},
		{"", ".*", true},
		{"", "c*", true},
		{"", "c*b*", true},
		{"b", "c*b*", true},
		{"ac", "ab*c", true},
		{"", "a", false},
		{"a", "", false},
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*ip*.", true},
		{"aaca", "ab*a*c*a", true},
		{"a", "ab*", true},
		{"a", "ab*c*", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
