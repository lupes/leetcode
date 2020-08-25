package question_1011_1020

import (
	"testing"
)

func Test_queryString(t *testing.T) {
	tests := []struct {
		S    string
		N    int
		want bool
	}{
		{"0110", 3, true},
		{"0110", 4, false},
		{"011011111", 1000, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := queryString(tt.S, tt.N); got != tt.want {
				t.Errorf("queryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
