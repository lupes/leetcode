package question_381_390

import "testing"

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{"test", "a", "b", false},
		{"test", "aa", "ab", false},
		{"test", "aa", "aab", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
