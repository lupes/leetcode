package question_71_80

func sortColors(nums []int) {
	var red, write, blue int
	for _, n := range nums {
		switch n {
		case 0:
			red++
		case 1:
			write++
		case 2:
			blue++
		}
	}
	for i := 0; i < red; i++ {
		nums[i] = 0
	}
	for i := 0; i < write; i++ {
		nums[i+red] = 1
	}
	for i := 0; i < blue; i++ {
		nums[i+red+write] = 2
	}
}
