package question_631_640

// 636. 函数的独占时间
// https://leetcode-cn.com/problems/exclusive-time-of-functions
// Topics: 栈

type Fun struct {
	Id   int
	Type bool
	Time int
}

func New(log string) Fun {
	var f, n = Fun{}, 0
	for _, c := range log {
		switch {
		case c == ':':
			n++
		case n == 0:
			f.Id = f.Id*10 + int(c) - '0'
		case n == 1 && c == 't':
			f.Type = true
		case n == 2:
			f.Time = f.Time*10 + int(c) - '0'
		}
	}
	return f
}

func exclusiveTime(n int, logs []string) []int {
	var res = make([]int, n)
	var stack []Fun
	for _, log := range logs {
		fun := New(log)
		if len(stack) > 0 && fun.Id == stack[len(stack)-1].Id && fun.Type != stack[len(stack)-1].Type {
			t := fun.Time - stack[len(stack)-1].Time + 1
			res[fun.Id] += t
			stack = stack[:len(stack)-1]
			for i := 0; i < len(stack); i++ {
				stack[i].Time += t
			}
		} else {
			stack = append(stack, fun)
		}
	}
	return res
}
