package question_791_800

import (
	"testing"
)

func Test_customSortString(t *testing.T) {
	tests := []struct {
		S    string
		T    string
		want string
	}{
		{"cbaz", "abcde", "cbade"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := customSortString(tt.S, tt.T); got != tt.want {
				t.Errorf("customSortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
