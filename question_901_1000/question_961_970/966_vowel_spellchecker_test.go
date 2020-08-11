package question_961_970

import (
	"reflect"
	"testing"
)

func Test_spellchecker(t *testing.T) {
	tests := []struct {
		wordlist []string
		queries  []string
		want     []string
	}{
		{[]string{"yellow"}, []string{"YellOw"}, []string{"yellow"}},
		{
			[]string{"KiTe", "kite", "hare", "Hare"},
			[]string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			[]string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := spellchecker(tt.wordlist, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spellchecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
