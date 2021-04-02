package question_1571_1580

import (
	"testing"
)

func Test_modifyString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"?zs", "azs"},
		{"b?a", "bca"},
		{"ubv?w", "ubvxw"},
		{"j?qg??b", "jkqghib"},
		{"??yw?ipkj?", "abywxipkjk"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := modifyString(tt.s); got != tt.want {
				t.Errorf("modifyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
