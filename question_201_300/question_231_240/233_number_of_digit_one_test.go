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
		{13},
		{231},
		{232},
		{1232},
		{21232},
		{91232},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			want := 0
			for i := 1; i <= tt.n; i++ {
				if strings.Count(strconv.Itoa(i), "1") > 0 {
					want++
				}
			}
			t.Logf("%d %d\n", tt.n, want)
			if got := countDigitOne(tt.n); got != want {
				t.Errorf("countDigitOne() = %v, want %v", got, want)
			}
		})
	}
}
