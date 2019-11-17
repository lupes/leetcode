package question_221_230

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"test", []int{3, 2, 3}, []int{3}},
		{"test", []int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
