package question_381_390

import (
	"fmt"
	"testing"
)

func TestRandomizedCollection(t *testing.T) {
	got := Constructor1()
	fmt.Printf("%v\n", got.Insert(1))
	fmt.Printf("%v\n", got.Insert(1))
	fmt.Printf("%v\n", got.Insert(2))
	fmt.Printf("%v\n", got.GetRandom())
	fmt.Printf("%v\n", got.GetRandom())
	fmt.Printf("%v\n", got.GetRandom())
	fmt.Printf("%v\n", got.GetRandom())
	fmt.Printf("%v\n", got.GetRandom())
	fmt.Printf("%v\n", got.GetRandom())

	fmt.Printf("%v\n", got.Remove(2))
	fmt.Printf("%v\n", got.Remove(1))
	fmt.Printf("%v\n", got.GetRandom())
	fmt.Printf("%v\n", got.GetRandom())
}
