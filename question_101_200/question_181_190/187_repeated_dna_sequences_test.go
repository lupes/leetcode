package question_181_190

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"test", "AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		{"test", "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
