package question_921_930

import (
	"testing"
)

func Test_minFlipsMonoIncr(t *testing.T) {
	tests := []struct {
		S    string
		want int
	}{
		{"00110", 1},
		{"010110", 2},
		{"00011000", 2},
		{"11011", 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minFlipsMonoIncr(tt.S); got != tt.want {
				t.Errorf("minFlipsMonoIncr() = %v, want %v", got, tt.want)
			}
		})
	}
}
