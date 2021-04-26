package question_1651_1660

import (
	"testing"
)

func Test_minimumDeletions(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aababbab", 2},
		{"bbaaaaabb", 2},
		{"aabbaaaaabb", 2},
		{"bbaabbaaaaabb", 4},
		{"aabbabbbbbaaabaabbbbb", 6},
		{"aabbbbaabababbbbaaaaaabbababaaabaabaabbbabbbbabbabbababaabaababbbbaaaaabbabbabaaaabbbabaaaabbaaabbbaabbaaaaabaa", 52},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minimumDeletions(tt.s); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
