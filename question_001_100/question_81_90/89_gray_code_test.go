package question_81_90

import (
	"reflect"
	"testing"
)

func Test_grayCode(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"test", 0, []int{0}},
		{"test", 1, []int{0, 1}},
		{"test", 2, []int{0, 1, 3, 2}},
		{"test", 11, []int{0, 1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
