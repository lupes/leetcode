package question_1451_1460

import (
	"testing"
)

func Test_arrangeWords(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{"Leetcode is cool", "Is cool leetcode"},
		{"Keep calm and code on", "On and keep calm code"},
		{"To be or not to be", "To be or to be not"},
		{"Jtik hrzrvhbmk gbo cfhmiqwonj ojkew ufvledv bomoeqt ops jgi zdhvbpbb zczmepdfpm jry rjazc titttcu", "Gbo ops jgi jry jtik ojkew rjazc ufvledv bomoeqt titttcu zdhvbpbb hrzrvhbmk cfhmiqwonj zczmepdfpm"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := arrangeWords(tt.text); got != tt.want {
				t.Errorf("arrangeWords() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
