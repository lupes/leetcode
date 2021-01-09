package question_1311_1320

import (
	"testing"
)

func Test_minFlips(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		c    int
		want int
	}{
		{2, 6, 5, 3},
		{4, 2, 7, 1},
		{1, 2, 3, 0},
		{0, 0, 3, 2},
		{0, 0, 7, 3},
		{7, 0, 0, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minFlips(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
