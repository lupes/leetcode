package question_1291_1300

import (
	"reflect"
	"testing"
)

func Test_sequentialDigits(t *testing.T) {
	tests := []struct {
		low  int
		high int
		want []int
	}{
		{100, 300, []int{123, 234}},
		{123, 300, []int{123, 234}},
		{123, 234, []int{123, 234}},
		{178546104, 812704742, []int{}},
		{1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sequentialDigits(tt.low, tt.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sequentialDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
