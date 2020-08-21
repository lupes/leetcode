package question_1011_1020

import (
	"testing"
)

func Test_smallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		K    int
		want int
	}{
		{1, 1},
		{2, -1},
		{3, 3},
		{7, 3},
		{23, 23},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := smallestRepunitDivByK(tt.K); got != tt.want {
				t.Errorf("smallestRepunitDivByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
