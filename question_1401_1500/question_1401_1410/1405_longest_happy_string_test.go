package question_1401_1410

import (
	"testing"
)

func Test_longestDiverseString(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		c    int
		want string
	}{
		//{0, 0, 0, ""},
		{1, 0, 0, "a"},
		{1, 1, 0, "ab"},
		{1, 1, 1, "abc"},
		{1, 1, 7, "ccaccbcc"},
		{2, 2, 1, "ababc"},
		{7, 1, 0, "aabaa"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := longestDiverseString(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("longestDiverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
