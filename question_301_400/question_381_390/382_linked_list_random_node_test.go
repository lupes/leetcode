package question_381_390

import (
	"fmt"
	"testing"

	. "github.com/lupes/leetcode/common"
)

func TestConstructor(t *testing.T) {
	this := Constructor2(NewList(1, 2, 3))
	fmt.Printf("%d\n", this.GetRandom())
	fmt.Printf("%d\n", this.GetRandom())
	fmt.Printf("%d\n", this.GetRandom())
	fmt.Printf("%d\n", this.GetRandom())
	fmt.Printf("%d\n", this.GetRandom())
	fmt.Printf("%d\n", this.GetRandom())
	fmt.Printf("%d\n", this.GetRandom())
	fmt.Printf("%d\n", this.GetRandom())
}
