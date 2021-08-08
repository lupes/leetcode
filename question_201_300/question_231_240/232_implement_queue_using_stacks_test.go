package question_231_240

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	got := Constructor()
	got.Push(1)
	fmt.Println(got)
	got.Push(2)
	fmt.Println(got)
	fmt.Printf("%d\n", got.Peek())
	fmt.Printf("%d\n", got.Pop())
	fmt.Println(got)
	got.Push(3)
	fmt.Println(got)
	fmt.Printf("%v\n", got.Empty())
	fmt.Printf("%d\n", got.Pop())
	fmt.Printf("%v\n", got.Empty())
}
