package question_131_140

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		s    string
		want [][]string
	}{
		{"aab", [][]string{{"aa", "b"}, {"a", "a", "b"}}},
		{"aaba", [][]string{{"aa", "b"}, {"a", "a", "b"}, {"a", "aba"}}},
		{"aabac", [][]string{{"aa", "b"}, {"a", "a", "b"}, {"a", "aba"}}},
		{"aabacc", [][]string{{"aa", "b"}, {"a", "a", "b"}, {"a", "aba"}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := partition(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %#v, want %v", got, tt.want)
			}
		})
	}
}
func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		//{"a", true},
		{"aa", true},
		{"aba", true},
		{"ab", false},
		{"abba", true},
		{"abcba", true},
		{"abccba", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
