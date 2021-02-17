package question_1471_1480

import (
	"fmt"
	"testing"
)

func TestSubrectangleQueries(t *testing.T) {
	queries := Constructor6([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	fmt.Printf("%d\n", queries.GetValue(0, 2))
	queries.UpdateSubrectangle(0, 0, 3, 2, 5)
	fmt.Printf("%d\n", queries.GetValue(0, 2))
	fmt.Printf("%d\n", queries.GetValue(3, 1))
	queries.UpdateSubrectangle(3, 0, 3, 2, 10)
	fmt.Printf("%d\n", queries.GetValue(3, 1))
	fmt.Printf("%d\n", queries.GetValue(0, 2))
}
