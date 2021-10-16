package question_281_290

// 282. 给表达式添加运算符
// https://leetcode-cn.com/problems/expression-add-operators
// Topics: 数学 字符串 回溯

func addOperators(num string, target int) []string {
	return addOperatorsHelper(num[:1], num[1:], target)
}

func addOperatorsHelper(prefix, num string, target int) []string {
	if len(num) == 0 {
		if operator(prefix) == target {
			return []string{prefix}
		}
		return nil
	}
	var res []string
	for _, c := range []string{"*", "+", "-", ""} {
		if c == "" && ((len(prefix) == 1 && prefix[0] == '0') ||
			(len(prefix) > 1 && (prefix[len(prefix)-1] == '0' && prefix[len(prefix)-2] < '0'))) {
			continue
		}
		tmp := prefix + c + num[:1]
		t := addOperatorsHelper(tmp, num[1:], target)
		res = append(res, t...)
	}
	return res
}

func operator(num string) int {
	num += "+0"
	var nums []int64
	var flag []int32
	for _, n := range num {
		if n >= '0' && n <= '9' {
			t := int64(n - '0')
			if len(nums) == len(flag) {
				nums = append(nums, t)
			} else {
				nums[len(nums)-1] = nums[len(nums)-1]*10 + t
			}
		} else {
			if len(flag) > 0 && flag[len(flag)-1] == '*' {
				nums[len(nums)-2] *= nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				flag = flag[:len(flag)-1]
			}
			flag = append(flag, n)
		}
	}
	var res = nums[0]
	for i := 1; i < len(nums); i++ {
		switch flag[i-1] {
		case '+':
			res += nums[i]
		case '-':
			res -= nums[i]
		}
	}
	return int(res)
}
