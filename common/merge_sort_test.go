package common

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 7, 8, 2, 4, 3, 0, 9, 5, 6}, []int{0, 1, 2, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mergeSort(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resortList(t *testing.T) {
	tests := []struct {
		root *ListNode
		want *ListNode
	}{
		{NewList(0, 1, 2, 3, 4, 5, 6), NewList(6, 5, 4, 3, 2, 1, 0)},
		{NewList(0, 1, 2, 3, 4, 5), NewList(5, 4, 3, 2, 1, 0)},
		{NewList(0, 1, 2, 3, 4), NewList(4, 3, 2, 1, 0)},
		{NewList(0, 1, 2, 3), NewList(3, 2, 1, 0)},
		{NewList(0, 1, 2), NewList(2, 1, 0)},
		{NewList(0, 1), NewList(1, 0)},
		{NewList(0), NewList(0)},
		{nil, nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := resortList(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resortList() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("%s\n", got)
			}
		})
	}
}
