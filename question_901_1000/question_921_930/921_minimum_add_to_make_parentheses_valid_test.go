package question_921_930

import (
	"testing"
)

func Test_minAddToMakeValid(t *testing.T) {
	tests := []struct {
		S    string
		want int
	}{
		{"())", 1},
		{"(((", 3},
		{"()", 0},
		{"()))((", 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minAddToMakeValid(tt.S); got != tt.want {
				t.Errorf("minAddToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
