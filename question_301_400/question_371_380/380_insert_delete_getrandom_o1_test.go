package question_371_380

import (
	"fmt"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	this := Constructor()
	fmt.Printf("%v\n", this.Insert(1))
	fmt.Printf("%v\n", this.Remove(2))
	fmt.Printf("%v\n", this.Insert(2))
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.Remove(1))
	fmt.Printf("%v\n", this.Insert(2))
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
	fmt.Printf("%v\n", this.GetRandom())
}
