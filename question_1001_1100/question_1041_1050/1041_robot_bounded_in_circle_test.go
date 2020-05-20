package question_1041_1050

import (
	"testing"
)

func Test_isRobotBounded(t *testing.T) {
	tests := []struct {
		instructions string
		want         bool
	}{
		{"GGLLGG", true},
		{"GG", false},
		{"GGLL", true},
		{"GLGLGGLGL", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isRobotBounded(tt.instructions); got != tt.want {
				t.Errorf("isRobotBounded() = %v, want %v", got, tt.want)
			}
		})
	}
}
