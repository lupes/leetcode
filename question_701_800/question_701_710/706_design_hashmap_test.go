package question_701_710

import (
	"fmt"
	"testing"
)

func TestConstructor4(t *testing.T) {
	got := Constructor4()
	got.Put(1000000, 100000)
	got.Put(1, 1)
	got.Put(2, 2)
	fmt.Printf("%d\n", got.Get(1000000))
	fmt.Printf("%d\n", got.Get(1))
	fmt.Printf("%d\n", got.Get(3))
	got.Put(2, 1)
	fmt.Printf("%d\n", got.Get(2))
	got.Remove(2)
	fmt.Printf("%d\n", got.Get(2))
}
