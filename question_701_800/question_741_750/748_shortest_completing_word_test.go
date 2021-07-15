package question_741_750

import (
	"testing"
)

func Test_shortestCompletingWord(t *testing.T) {
	tests := []struct {
		licensePlate string
		words        []string
		want         string
	}{
		{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}, "steps"},
		{"1s3 456", []string{"looks", "pest", "stew", "show"}, "pest"},
		{"Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}, "husband"},
		{"OgEu755", []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}, "enough"},
		{"iMSlpe4", []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"}, "simple"},
		{"Ar16259", []string{"nature", "though", "party", "food", "any", "democratic", "building", "eat", "structure", "our"}, "party"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shortestCompletingWord(tt.licensePlate, tt.words); got != tt.want {
				t.Errorf("shortestCompletingWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
