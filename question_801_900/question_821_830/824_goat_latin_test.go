package question_821_830

import (
	"testing"
)

func Test_toGoatLatin(t *testing.T) {
	tests := []struct {
		S    string
		want string
	}{
		{"I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{"The quick brown fox jumped over the lazy dog", "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := toGoatLatin(tt.S); got != tt.want {
				t.Errorf("toGoatLatin() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
