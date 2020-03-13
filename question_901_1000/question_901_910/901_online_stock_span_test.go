package question_901_910

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	got := Constructor()
	fmt.Println(got.Next(100))
	fmt.Println(got.Next(80))
	fmt.Println(got.Next(60))
	fmt.Println(got.Next(70))
	fmt.Println(got.Next(60))
	fmt.Println(got.Next(75))
	fmt.Println(got.Next(85))
}
