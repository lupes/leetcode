package question_431_440

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
		{"a", "b", []int{}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findAnagrams(tt.s, tt.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
