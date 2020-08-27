package question_1161_1170

import (
	"reflect"
	"testing"
)

func Test_numSmallerByFrequency(t *testing.T) {
	tests := []struct {
		queries []string
		words   []string
		want    []int
	}{
		{[]string{"cbd"}, []string{"zaaaz"}, []int{1}},
		{[]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}, []int{1, 2}},
		{
			[]string{"bba", "abaaaaaa", "aaaaaa", "bbabbabaab", "aba", "aa", "baab", "bbbbbb", "aab", "bbabbaabb"},
			[]string{"aaabbb", "aab", "babbab", "babbbb", "b", "bbbbbbbbab", "a", "bbbbbbbbbb", "baaabbaab", "aa"},
			[]int{6, 1, 1, 2, 3, 3, 3, 1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numSmallerByFrequency(tt.queries, tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSmallerByFrequency() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
