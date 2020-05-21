package question_1071_1080

import (
	"reflect"
	"testing"
)

func Test_findOcurrences(t *testing.T) {
	tests := []struct {
		text   string
		first  string
		second string
		want   []string
	}{
		{"alice is a good girl she is a good student", "a", "good", []string{"girl", "student"}},
		{"we will we will rock you", "we", "will", []string{"we", "rock"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findOcurrences(tt.text, tt.first, tt.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOcurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
