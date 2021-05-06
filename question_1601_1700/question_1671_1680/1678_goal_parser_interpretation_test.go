package question_1671_1680

import (
	"testing"
)

func Test_interpret(t *testing.T) {
	tests := []struct {
		command string
		want    string
	}{
		{"G()(al)", "Goal"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := interpret(tt.command); got != tt.want {
				t.Errorf("interpret() = %v, want %v", got, tt.want)
			}
		})
	}
}
