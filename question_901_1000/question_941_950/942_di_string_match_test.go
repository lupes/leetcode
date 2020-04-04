package question_941_950

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	tests := []struct {
		S    string
		want []int
	}{
		{"IDID", []int{0, 4, 1, 3, 2}},
		{"III", []int{0, 1, 2, 3}},
		{"DDI", []int{3, 2, 0, 1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := diStringMatch(tt.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
