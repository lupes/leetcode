package question_221_230

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{"test", []int{0}, []string{"0"}},
		{"test", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"test", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
