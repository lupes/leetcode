package question_341_350

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"test", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"test", []int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
