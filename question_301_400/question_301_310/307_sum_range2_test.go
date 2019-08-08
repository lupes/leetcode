package question_301_310

import (
	"fmt"
	"testing"
)

func TestNumArray_SumRange2(t *testing.T) {

	self := Constructor2([]int{1, 3, 5})
	fmt.Printf("%v\n", self)

	got := self.SumRange(0, 2)
	if 9 != got {
		t.Errorf("got=%v want=%v", got, 9)
	}

	self.Update(1, 2)
	fmt.Printf("%v\n", self)

	got = self.SumRange(0, 2)
	if 8 != got {
		t.Errorf("got=%v want=%v", got, 8)
	}

}
