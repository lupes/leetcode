package question_201_210

// 210. 课程表 II
// https://leetcode-cn.com/problems/course-schedule-ii
// Topics: 深度优先搜索 广度优先搜索 图 拓扑排序

func findOrder(numCourses int, prerequisites [][]int) []int {
	var tmp = make(map[int][]int, len(prerequisites))
	for _, n := range prerequisites {
		tmp[n[0]] = append(tmp[n[0]], n[1])
	}

	var flag = make(map[int]struct{}, numCourses)
	var res = make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if _, ok := tmp[i]; !ok {
			flag[i] = struct{}{}
			res = append(res, i)
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
				res = append(res, n)
				change = true
				delete(tmp, n)
			}
		}
	}

	if len(res) == numCourses {
		return res
	}
	return nil
}
