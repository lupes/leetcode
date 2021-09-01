package question_341_350

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		nestedList []*NestedInteger
	}{
		{
			[]*NestedInteger{
				{flag: false, lists: []*NestedInteger{{flag: true, value: 1}, {flag: true, value: 1}}},
				{flag: true, value: 2},
				{flag: false, lists: []*NestedInteger{{flag: true, value: 1}, {flag: true, value: 1}}},
			},
		},
		{
			[]*NestedInteger{
				{flag: true, value: 1},
				{flag: false, lists: []*NestedInteger{{flag: true, value: 4}, {flag: false, lists: []*NestedInteger{{flag: true, value: 6}}}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := Constructor(tt.nestedList)
			fmt.Printf("%v\n", got.nums)
		})
	}
}
