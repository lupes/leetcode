package question_61_70

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"test#1", []int{1}, []int{2}},
		{"test#2", []int{1, 2}, []int{1, 3}},
		{"test#3", []int{1, 9}, []int{2, 0}},
		{"test#4", []int{9, 9}, []int{1, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
