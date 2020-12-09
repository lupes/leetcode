package question_0011_0020

import "testing"

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"", "", 0},
		{"", "l", -1},
		{"l", "", 0},
		{"l", "ll", -1},
		{"l", "l", 0},
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
