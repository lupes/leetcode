package question_671_680

import (
	"fmt"
	"testing"
)

func TestMapSum(t *testing.T) {
	got := Constructor()
	got.Insert("apple", 3)
	fmt.Printf("%d\n", got.Sum("ap"))
	got.Insert("app", 2)
	fmt.Printf("%d\n", got.Sum("ap"))
}
