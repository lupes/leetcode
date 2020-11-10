package question_1361_1370

import (
	"testing"
)

func Test_sortString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"aaaabbbbcccc", "abccbaabccba"},
		{"rat", "art"},
		{"leetcode", "cdelotee"},
		{"ggggggg", "ggggggg"},
		{"spo", "ops"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortString(tt.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
