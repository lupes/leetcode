package question_1021_1030

import (
	"sort"
)

// 1024. 视频拼接
// https://leetcode-cn.com/problems/video-stitching
// Topics: 动态规划

func videoStitching(clips [][]int, T int) int {
	var res = make([]int, T+1)
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})
	for _, clip := range clips {
		if clip[0] >= T {
			break
		}
		tmp := res[clip[0]] + 1
		if clip[0] != 0 && tmp == 1 {
			return -1
		}
		for i := clip[0] + 1; i <= clip[1] && i <= T; i++ {
			if res[i] == 0 || res[i] > tmp {
				res[i] = tmp
			}
		}
	}
	if res[T] == 0 {
		return -1
	}
	return res[T]
}
