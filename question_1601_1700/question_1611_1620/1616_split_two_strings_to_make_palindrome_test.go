package question_1611_1620

import (
	"testing"
)

func Test_checkPalindromeFormation(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"x", "y", true},
		{"aba", "xyz", true},
		{"fecab", "abdef", true},
		{"abdef", "fecab", true},
		{"ulacfd", "jizalu", true},
		{"abcdcbe", "abcdefa", true},
		{"fgixkdymrtiqifbmwjhuwdukaqfjekzckyjyxfbdiswmmwsidbfxyjykczkejfqakudwuhjwmbfizciijajfncqy", "yqcnfjajiiczbzgcziiriyhfxlqrdtijusntxzazzxncqnuyryrognehxsgkmkzgtmtpdpkffxqvgtmcifissukj", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkPalindromeFormation(tt.a, tt.b); got != tt.want {
				t.Errorf("checkPalindromeFormation() = %v, want %v", got, tt.want)
			}
		})
	}
}

// fgixkdymrtiqifbmwjhuwdukaqfjekzckyjyxfbdiswmmwsidbfxyjykczkejfqakudwuhjwmbfizciijaj
// jajiiczbzgcziiriyhfxlqrdtijusntxzazzxncqnuyryrognehxsgkmkzgtmtpdpkffxqvgtmcifissukj
