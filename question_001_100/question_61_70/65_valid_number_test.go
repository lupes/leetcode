package question_61_70

import "testing"

func Test_isNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"test", "0", true},
		{"test", ".", false},
		{"test", "+.", false},
		{"test", "+0", true},
		{"test", "0.1", true},
		{"test", "abc", false},
		{"test", "1 a", false},
		{"test", "2e10", true},
		{"test", " -90e3  ", true},
		{"test", " 1e ", false},
		{"test", "e3", false},
		{"test", " 6e-1 ", true},
		{"test", " 99e2.5 ", false},
		{"test", " 53.5e93 ", true},
		{"test", " --6 ", false},
		{"test", " -+3 ", false},
		{"test", " ++3 ", false},
		{"test", " 95a54e53 ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.s); got != tt.want {
				t.Errorf("isNumber(%#v) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
