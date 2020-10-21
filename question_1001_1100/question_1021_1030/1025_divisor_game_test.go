package question_1021_1030

import (
	"testing"
)

func Test_divisorGame(t *testing.T) {
	tests := []struct {
		N    int
		want bool
	}{
		{2, true},
		{3, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := divisorGame(tt.N); got != tt.want {
				t.Errorf("divisorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
