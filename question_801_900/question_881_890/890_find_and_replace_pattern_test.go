package question_881_890

import (
	"reflect"
	"testing"
)

func Test_findAndReplacePattern(t *testing.T) {
	tests := []struct {
		words   []string
		pattern string
		want    []string
	}{
		{[]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb", []string{"mee", "aqq"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findAndReplacePattern(tt.words, tt.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAndReplacePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
