package offer

import (
	"testing"
)

func Test_replaceSpace(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"We are happy.", "We%20are%20happy."},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := replaceSpace(tt.s); got != tt.want {
				t.Errorf("replaceSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
