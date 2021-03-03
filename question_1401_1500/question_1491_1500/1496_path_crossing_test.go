package question_1491_1500

import (
	"testing"
)

func Test_isPathCrossing(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"NES", false},
		{"NESWW", true},
		{"NNSWWEWSSESSWENNW", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isPathCrossing(tt.path); got != tt.want {
				t.Errorf("isPathCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
