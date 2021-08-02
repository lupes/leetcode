package question_201_210

// 207. 课程表
// https://leetcode-cn.com/problems/course-schedule
// Topics: 深度优先搜索 广度优先搜索 图 拓扑排序

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) < 2 {
		return true
	}
	var tmp = make(map[int][]int)
	for _, n := range prerequisites {
		tmp[n[0]] = append(tmp[n[0]], n[1])
	}

	var flag = make(map[int]struct{})
	for i := 0; i < numCourses; i++ {
		if _, ok := tmp[i]; !ok {
			flag[i] = struct{}{}
		}
	}

	for change := true; change; {
		change = false
		for n, ns := range tmp {
			all := true
			for _, t := range ns {
				if _, ok := flag[t]; !ok {
					all = false
				}
			}
			if all {
				flag[n] = struct{}{}
				change = true
				delete(tmp, n)
			}
		}
	}

	return len(tmp) == 0
}
