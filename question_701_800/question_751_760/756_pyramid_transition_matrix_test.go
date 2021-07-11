package question_751_760

import (
	"testing"
)

func Test_pyramidTransition(t *testing.T) {
	tests := []struct {
		bottom  string
		allowed []string
		want    bool
	}{
		{"BCD", []string{"BCG", "CDE", "GEA", "FFF"}, true},
		{"AABA", []string{"AAA", "AAB", "ABA", "ABB", "BAC"}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pyramidTransition(tt.bottom, tt.allowed); got != tt.want {
				t.Errorf("pyramidTransition() = %v, want %v", got, tt.want)
			}
		})
	}
}
