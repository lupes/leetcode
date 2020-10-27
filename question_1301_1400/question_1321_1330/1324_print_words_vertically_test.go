package question_1321_1330

import (
	"reflect"
	"testing"
)

func Test_printVertically(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"HOW ARE YOU", []string{"HAY", "ORO", "WEU"}},
		{"TO BE OR NOT TO BE", []string{"TBONTB", "OEROOE", "   T"}},
		{"CONTEST IS COMING", []string{"CIC", "OSO", "N M", "T I", "E N", "S G", "T"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := printVertically(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}
