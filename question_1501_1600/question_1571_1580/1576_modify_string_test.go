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
		{"ubv?w", "ubvaw"},
		{"j?qg??b", "jaqgacb"},
		{"??yw?ipkj?", "abywaipkja"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := modifyString(tt.s); got != tt.want {
				t.Errorf("modifyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
