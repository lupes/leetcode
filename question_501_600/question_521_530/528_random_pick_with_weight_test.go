package question_521_530

import (
	"fmt"
	"testing"
)

func TestSolution_PickIndex(t *testing.T) {

	tests := []struct {
		w []int
	}{
		{[]int{1}},
		{[]int{1, 3}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			this := Constructor(tt.w)
			fmt.Printf("%+v\n", this)
			for i := 0; i < 9; i++ {
				fmt.Printf("index:%d\n", this.PickIndex())
			}
		})
	}
}
