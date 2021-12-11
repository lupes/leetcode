package question_911_920

// 911. 在线选举
// https://leetcode-cn.com/problems/online-election
// Topics: 二分查找

type t struct {
	time int
	id   int
}

type TopVotedCandidate struct {
	ts []t
	l  int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var flag = make([]t, len(times))
	var m = make(map[int]int)
	var max, id = 0, 0
	for i, p := range persons {
		m[p]++
		if m[p] >= max {
			max = m[p]
			id = p
		}
		flag[i] = t{time: times[i], id: id}
	}
	flag = append(flag, t{time: 1e9 + 1, id: id})
	return TopVotedCandidate{
		ts: flag,
		l:  len(times),
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	left, right := 0, this.l
	for right > left {
		mid := (right+left)/2 + 1
		if this.ts[mid].time == t {
			return this.ts[mid].id
		} else if this.ts[mid].time > t {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return this.ts[left].id
}
