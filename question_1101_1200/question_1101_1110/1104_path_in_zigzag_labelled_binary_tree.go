package question_1101_1110

// 1104. 二叉树寻路
// https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree
// Topics: 树 数学

var tree = [][2]int{
	{1, 1}, {2, 3}, {4, 7}, {8, 15}, {16, 31},
	{32, 63}, {64, 127}, {128, 255}, {256, 511}, {512, 1023},
	{1024, 2047}, {2048, 4095}, {4096, 8191}, {8192, 16383}, {16384, 32767},
	{32768, 65535}, {65536, 131071}, {131072, 262143}, {262144, 524287}, {524288, 1048575},
}

func pathInZigZagTree(label int) []int {
	var i, f = 0, [2]int{0, 0}
	for i, f = range tree {
		if f[1] >= label {
			break
		}
	}
	var res = make([]int, i+1)
	for ; i > 0; i-- {
		res[i] = label
		if i%2 == 0 {
			label = tree[i-1][1] - (label-tree[i][0])/2
		} else {
			label = tree[i-1][0] + (tree[i][1]-label)/2
		}
	}
	res[0] = 1
	return res
}
