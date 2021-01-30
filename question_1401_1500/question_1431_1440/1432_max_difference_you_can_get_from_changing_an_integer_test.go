package question_1431_1440

import (
	"testing"
)

func Test_maxDiff(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{555, 888},
		{123456, 820000},
		{9, 8},
		{10000, 80000},
		{9288, 8700},
		{111, 888},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxDiff(tt.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
