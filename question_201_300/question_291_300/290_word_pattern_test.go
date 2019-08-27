package question_291_300

import "testing"

func Test_wordPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		str     string
		want    bool
	}{
		{"test", "abba", "dog cat cat dog", true},
		{"test", "abba", "dog cat cat fish", false},
		{"test", "aaaa", "dog cat cat dog", false},
		{"test", "abba", "dog dog dog dog", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.pattern, tt.str); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
