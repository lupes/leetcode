package question_851_860

import (
	"testing"
)

func Test_buddyStrings(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want bool
	}{
		{"ab", "ba", true},
		{"ab", "ab", false},
		{"aa", "aa", true},
		{"aaaaaaabc", "aaaaaaacb", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := buddyStrings(tt.A, tt.B); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
