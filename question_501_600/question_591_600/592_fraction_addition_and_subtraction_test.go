package question_591_600

import (
	"testing"
)

func Test_fractionAddition(t *testing.T) {
	tests := []struct {
		expression string
		want       string
	}{
		{"-1/2+1/2", "0/1"},
		{"-1/2+1/2-1/10", "0/1"},
		{"-1/2+1/2+1/3", "1/3"},
		{"5/3+1/3", "2/1"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := fractionAddition(tt.expression); got != tt.want {
				t.Errorf("fractionAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
