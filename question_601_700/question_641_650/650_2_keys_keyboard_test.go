package question_641_650

import (
	"testing"
)

func Test_minSteps(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		//{1, 0},
		//{2, 2},
		//{3, 3},
		//{4, 4},
		//{5, 5},
		//{6, 5},
		//{7, 7},
		//{8, 6},
		//{9, 6},
		//{10, 7},
		//{18, 8},
		{100, 14},
		{1000, 131},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minSteps(tt.n); got != tt.want {
				t.Errorf("minSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
