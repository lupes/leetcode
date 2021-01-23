package question_1411_1420

import (
	"testing"
)

func Test_getHappyString(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want string
	}{
		{1, 3, "c"},
		{1, 4, ""},
		{3, 9, "cab"},
		{10, 100, "abacbabacb"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getHappyString(tt.n, tt.k); got != tt.want {
				t.Errorf("getHappyString() = %v, want %v", got, tt.want)
			}
		})
	}
}
