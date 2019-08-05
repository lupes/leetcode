package question_341_350

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{"test", []byte("hello"), []byte("olleh")},
		{"test", []byte("Hannah"), []byte("hannaH")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("reverseString got=%s want=%s", string(tt.s), string(tt.want))
			}
		})
	}
}
