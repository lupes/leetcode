package question_1071_1080

import (
	"testing"
)

func Test_numTilePossibilities(t *testing.T) {
	tests := []struct {
		tiles string
		want  int
	}{
		{"AAB", 8},
		{"AAABBC", 188},
		{"ABCDEFG", 13699},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numTilePossibilities(tt.tiles); got != tt.want {
				t.Errorf("numTilePossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
