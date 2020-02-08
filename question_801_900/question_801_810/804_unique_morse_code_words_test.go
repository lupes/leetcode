package question_801_810

import (
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := uniqueMorseRepresentations(tt.words); got != tt.want {
				t.Errorf("uniqueMorseRepresentations() = %v, want %v", got, tt.want)
			}
		})
	}
}
