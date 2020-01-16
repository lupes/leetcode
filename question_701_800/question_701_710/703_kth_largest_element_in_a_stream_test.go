package question_701_710

import (
	"fmt"
	"testing"
)

func TestConstructor2(t *testing.T) {
	this := Constructor2(3, []int{4, 5, 8, 2})
	fmt.Printf("%d\n", this.Add(3))
	fmt.Printf("%d\n", this.Add(5))
	fmt.Printf("%d\n", this.Add(10))
	fmt.Printf("%d\n", this.Add(9))
	fmt.Printf("%d\n", this.Add(4))

	this = Constructor2(5, []int{4, 5, 8, 2})
	fmt.Printf("%d\n", this.Add(3))
}
