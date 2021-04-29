package question_1661_1670

import (
	"testing"
)

func Test_getSmallestString(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want string
	}{
		{3, 27, "aay"},
		{5, 73, "aaszz"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getSmallestString(tt.n, tt.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
