package question_491_500

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{"test", []string{"A"}, []string{"A"}},
		{"test", []string{"hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
		{"test", []string{"Aasdfghjkl", "Qwertyuiop", "zZxcvbnm"}, []string{"Aasdfghjkl", "Qwertyuiop", "zZxcvbnm"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
