package question_781_790

import (
	"reflect"
	"sort"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	tests := []struct {
		S    string
		want []string
	}{
		{"1", []string{"1"}},
		{"a", []string{"a", "A"}},
		{"a1", []string{"a1", "A1"}},
		{"a1a", []string{"a1a", "A1a", "a1A", "A1A"}},
		{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		{"12345", []string{"12345"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := letterCasePermutation(tt.S)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
