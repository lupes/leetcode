package question_881_890

import (
	"reflect"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	tests := []struct {
		A    string
		B    string
		want []string
	}{
		{"this apple is sweet", "this apple is sour", []string{"sweet", "sour"}},
		{" apple apple", "banana", []string{"banana"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := uncommonFromSentences(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}
