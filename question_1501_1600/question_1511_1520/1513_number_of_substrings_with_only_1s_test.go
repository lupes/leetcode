package question_1511_1520

import (
	"testing"
)

func Test_numSub(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"0110111", 9},
		{"101", 2},
		{"111111", 21},
		{"000", 0},
		{"011", 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numSub(tt.s); got != tt.want {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
