package question_911_920

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	tests := []struct {
		A    []string
		B    []string
		want []string
	}{
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"e", "o"},
			[]string{"facebook", "google", "leetcode"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"l", "e"},
			[]string{"apple", "google", "leetcode"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"e", "oo"},
			[]string{"facebook", "google"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"lo", "eo"},
			[]string{"google", "leetcode"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"ec", "oc", "ceo"},
			[]string{"facebook", "leetcode"},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := wordSubsets(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
