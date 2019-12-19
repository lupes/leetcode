package question_1101_1110

import (
	"testing"
)

func Test_parseBoolExpr(t *testing.T) {
	tests := []struct {
		expression string
		want       bool
	}{
		{"t", true},
		{"f", false},
		{"!(t)", false},
		{"!(f)", true},
		{"|(f,t)", true},
		{"&(f,t)", false},
		{"|(&(t,f,t),!(t))", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := parseBoolExpr(tt.expression); got != tt.want {
				t.Errorf("parseBoolExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
