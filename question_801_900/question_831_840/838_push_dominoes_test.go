package question_831_840

import (
	"testing"
)

func Test_pushDominoes(t *testing.T) {
	tests := []struct {
		dominoes string
		want     string
	}{
		{".L.R...LR..L..", "LL.RR.LLRRLL.."},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pushDominoes(tt.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
