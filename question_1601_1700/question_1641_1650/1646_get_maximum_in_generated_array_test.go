package question_1641_1650

import (
	"testing"
)

func Test_getMaximumGenerated(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{7, 3},
		{2, 1},
		{3, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getMaximumGenerated(tt.n); got != tt.want {
				t.Errorf("getMaximumGenerated() = %v, want %v", got, tt.want)
			}
		})
	}
}
