package question_841_850

import (
	"testing"
)

func Test_backspaceCompare(t *testing.T) {
	tests := []struct {
		S    string
		T    string
		want bool
	}{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a##c", "#a#c", true},
		{"a#c", "b", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := backspaceCompare(tt.S, tt.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
