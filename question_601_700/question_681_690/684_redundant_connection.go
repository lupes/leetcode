package question_681_690

// 684. 冗余连接
// https://leetcode-cn.com/problems/redundant-connection
// Topics: 树 并查集 图

func findRedundantConnection(edges [][]int) []int {
	var flag = make([]map[int]struct{}, 0)
	for _, edge := range edges {
		var ok1, ok2 bool
		var i1, i2 int
		for i, t := range flag {
			if _, ok := t[edge[0]]; ok {
				ok1, i1 = ok, i
			}
			if _, ok := t[edge[1]]; ok {
				ok2, i2 = ok, i
			}
		}
		if !ok1 && !ok2 {
			tmp := make(map[int]struct{})
			tmp[edge[0]] = struct{}{}
			tmp[edge[1]] = struct{}{}
			flag = append(flag, tmp)
		} else if ok1 && !ok2 {
			flag[i1][edge[1]] = struct{}{}
		} else if !ok1 && ok2 {
			flag[i2][edge[0]] = struct{}{}
		} else {
			if i1 == i2 {
				return edge
			} else if i1 > i2 {
				i1, i2 = i2, i1
			}
			for k := range flag[i2] {
				flag[i1][k] = struct{}{}
			}
			for ; i2 < len(flag)-1; i2++ {
				flag[i2] = flag[i2+1]
			}
			flag = flag[:len(flag)-1]
		}
	}
	return nil
}
