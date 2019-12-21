package question_221_230

import (
	"testing"
)

func Test_calculate2(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"1-2", -1},
		{"2+2", 4},
		{"2*2", 4},
		{"3+2*2", 7},
		{"3+2*2*2", 11},
		{" 3/2 ", 1},
		{" 3+5 / 2 ", 5},
		{"100*100*100*100/10000", 10000},
		{"100*100*100*100/10000+1", 10001},
		{"1/2/3/4/5/6/7", 0},
		{"1/2/3/4/5/6/7+1", 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := calculate2(tt.s); got != tt.want {
				t.Errorf("calculate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
