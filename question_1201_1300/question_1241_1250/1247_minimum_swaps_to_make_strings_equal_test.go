package question_1241_1250

import (
	"testing"
)

func Test_minimumSwap(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{"xx", "yy", 1},
		{"xy", "yx", 2},
		{"xx", "xy", -1},
		{"xxyyxyxyxx", "xyyxyxxxyx", 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minimumSwap(tt.s1, tt.s2); got != tt.want {
				t.Errorf("minimumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
