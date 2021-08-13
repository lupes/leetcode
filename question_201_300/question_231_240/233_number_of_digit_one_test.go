package question_231_240

import (
	"strconv"
	"strings"
	"testing"
)

func Test_countDigitOne(t *testing.T) {
	tests := []struct {
		n int
	}{
		{1},
		{10},
		{13},
		{231},
		{232},
		{1232},
		{8794835},
		//{82488329},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			want := 0
			for i := 1; i <= tt.n; i++ {
				want += strings.Count(strconv.Itoa(i), "1")
			}
			if got := countDigitOne(tt.n); got != want {
				t.Errorf("countDigitOne() = %v, want %v", got, want)
			}
		})
	}
}
