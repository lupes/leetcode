package question_1331_1340

import (
	"testing"
)

func Test_angleClock(t *testing.T) {
	tests := []struct {
		hour    int
		minutes int
		want    float64
	}{
		{12, 30, 165},
		{3, 30, 75},
		{3, 15, 7.5},
		{4, 50, 155},
		{12, 0, 0},
		{9, 45, 22.5},
		{9, 50, 5},
		{11, 10, 85},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := angleClock(tt.hour, tt.minutes); got != tt.want {
				t.Errorf("angleClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
