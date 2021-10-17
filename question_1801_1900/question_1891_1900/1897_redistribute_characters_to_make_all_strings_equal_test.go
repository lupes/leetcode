package question_1891_1900

import (
	"testing"
)

func Test_makeEqual(t *testing.T) {
	tests := []struct {
		words []string
		want  bool
	}{
		{[]string{"abc", "aabc", "bc"}, true},
		{[]string{"ab", "a"}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := makeEqual(tt.words); got != tt.want {
				t.Errorf("makeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
