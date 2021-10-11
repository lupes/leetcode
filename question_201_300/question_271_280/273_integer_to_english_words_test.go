package question_271_280

import (
	"testing"
)

func Test_numberToWordsHelper(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{123, "One Hundred Twenty Three"},
		{119, "One Hundred Nineteen"},
		{109, "One Hundred Nine"},
		{223, "Two Hundred Twenty Three"},
		{981, "Nine Hundred Eighty One"},
		{980, "Nine Hundred Eighty"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberToWordsHelper(tt.num); got != tt.want {
				t.Errorf("numberToWordsHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberToWords(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{0, "Zero"},
		{100, "One Hundred"},
		{1001, "One Thousand One"},
		{1000, "One Thousand"},
		{50868, "Fifty Thousand Eight Hundred Sixty Eight"},
		{12345, "Twelve Thousand Three Hundred Forty Five"},
		{1000000, "One Million"},
		{1234567, "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"},
		{1234567891, "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberToWords(tt.num); got != tt.want {
				t.Errorf("numberToWords() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
