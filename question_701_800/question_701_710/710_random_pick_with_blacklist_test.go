package question_701_710

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {

	got := Constructor(3, []int{1})
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	fmt.Printf("%d\n", got.Pick())
	//Constructor(50, []int{1, 3, 8, 49, 27, 26})
}
