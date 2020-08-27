package question_1071_1080

import (
	"testing"
)

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "COOD", ""},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := gcdOfStrings(tt.str1, tt.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
