package question_01_10

import (
	"reflect"
	"testing"
)

var (
	test_1_l1   = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	test_1_l2   = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	test_1_want = &ListNode{7, &ListNode{0, &ListNode{8, nil}}}

	test_2_l1   = &ListNode{2, nil}
	test_2_l2   = &ListNode{5, nil}
	test_2_want = &ListNode{7, nil}

	test_3_l1   = &ListNode{5, nil}
	test_3_l2   = &ListNode{5, nil}
	test_3_want = &ListNode{0, &ListNode{1, nil}}

	test_4_l1   = &ListNode{5, nil}
	test_4_l2   = &ListNode{5, &ListNode{1, nil}}
	test_4_want = &ListNode{0, &ListNode{2, nil}}
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test#1", args{test_1_l1, test_1_l2}, test_1_want},
		{"test#2", args{test_2_l1, test_2_l2}, test_2_want},
		{"test#3", args{test_3_l1, test_3_l2}, test_3_want},
		{"test#4", args{test_4_l1, test_4_l2}, test_4_want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
