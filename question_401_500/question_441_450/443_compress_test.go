package question_441_450

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		name  string
		chars []byte
		want  []byte
	}{
		{"test", []byte{'a'}, []byte{'a'}},
		{"test", []byte{'a', 'b'}, []byte{'a', 'b'}},
		{"test", []byte{'a', 'a'}, []byte{'a', '2'}},
		{"test", []byte{'a', 'a', 'a'}, []byte{'a', '3'}},
		{"test", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}},
		{"test", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.chars); !reflect.DeepEqual(tt.chars[:got], tt.want) {
				t.Errorf("compress() = %v, want %v", tt.chars[:got], tt.want)
			}
		})
	}
}
