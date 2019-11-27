package question_651_660

import (
	"testing"
)

func Test_judgeCircle(t *testing.T) {
	tests := []struct {
		moves string
		want  bool
	}{
		{"UD", true},
		{"LL", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := judgeCircle(tt.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
