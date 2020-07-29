package question_631_640

import (
	"strconv"
	"strings"
)

// 636. 函数的独占时间
// https://leetcode-cn.com/problems/exclusive-time-of-functions
// Topics: 栈

type Fun struct {
	Id   int
	Type bool
	Time int
}

func New(log string) Fun {
	arr := strings.Split(log, ":")
	id, _ := strconv.Atoi(arr[0])
	t, _ := strconv.Atoi(arr[2])
	return Fun{
		Id:   id,
		Type: arr[1] == "start",
		Time: t,
	}
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
