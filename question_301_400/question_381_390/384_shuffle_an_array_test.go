package question_381_390

import (
	"fmt"
	"testing"
)

func TestSolution_Shuffle(t *testing.T) {
	this := Constructor([]int{1, 2, 3, 4, 5, 6})
	for i := 0; i < 10; i++ {
		fmt.Printf("%+v\n", this.Shuffle())
	}
}
