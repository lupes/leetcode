package question_851_860

import (
	"testing"
)

func Test_lemonadeChange(t *testing.T) {
	tests := []struct {
		bills []int
		want  bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5, 5, 10, 10, 20}, false},
		{[]int{5, 5, 5, 10, 10, 20}, true},
		{[]int{5, 5, 5, 20}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lemonadeChange(tt.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
