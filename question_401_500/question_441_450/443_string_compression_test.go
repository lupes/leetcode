package question_441_450

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	tests := []struct {
		chars []byte
		want  []byte
	}{
		{[]byte{'a'}, []byte{'a'}},
		{[]byte{'a', 'b'}, []byte{'a', 'b'}},
		{[]byte{'a', 'a'}, []byte{'a', '2'}},
		{[]byte{'a', 'a', 'a'}, []byte{'a', '3'}},
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := compress2(tt.chars); !reflect.DeepEqual(tt.chars[:got], tt.want) {
				t.Errorf("compress() = %v, want %v", tt.chars[:got], tt.want)
			}
		})
	}
}
