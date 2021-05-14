package question_1651_1660

import (
	"fmt"
	"testing"
)

func TestOrderedStream_Insert(t *testing.T) {
	orderedStream := Constructor(5)
	fmt.Printf("%v\n", orderedStream.Insert(3, "c"))
	fmt.Printf("%v\n", orderedStream.Insert(1, "a"))
	fmt.Printf("%v\n", orderedStream.Insert(2, "b"))
	fmt.Printf("%v\n", orderedStream.Insert(5, "e"))
	fmt.Printf("%v\n", orderedStream.Insert(4, "d"))
}
