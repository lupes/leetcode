package question_1041_1050

import (
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		S    string
		want string
	}{
		{"abbaca", "ca"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeDuplicates(tt.S); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
