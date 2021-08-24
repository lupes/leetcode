package question_781_790

import (
	"math"
)

// 787. K 站中转内最便宜的航班
// https://leetcode-cn.com/problems/cheapest-flights-within-k-stops
// Topics: 堆 广度优先搜索 动态规划

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, K+2)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	for i := 0; i <= K+1; i++ {
		dp[src][i] = 0
	}
	for k := 1; k <= K+1; k++ {
		for _, flight := range flights {
			if dp[flight[0]][k-1] != math.MaxInt64 {
				if dp[flight[0]][k-1]+flight[2] < dp[flight[1]][k] {
					dp[flight[1]][k] = dp[flight[0]][k-1] + flight[2]
				}
			}
		}
	}
	if dp[dst][K+1] == math.MaxInt64 {
		return -1
	}
	return dp[dst][K+1]
}

// 超时
func findCheapestPrice2(n int, flights [][]int, src int, dst int, K int) int {
	var flightMap = make(map[int][][]int, n)
	for _, flight := range flights {
		flightMap[flight[0]] = append(flightMap[flight[0]], flight)
	}
	flag := make(map[int]bool)
	flag[src] = true
	res := findCheapestPriceHelper(flightMap, flag, src, dst, 0, math.MaxInt64, K+1)
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

func findCheapestPriceHelper(flightMap map[int][][]int, flag map[int]bool, src, dst, price, min, K int) int {
	if K == 0 && src != dst {
		return math.MaxInt64
	} else if src == dst {
		return price
	}
	if price > min {
		return min
	}
	for _, flights := range flightMap[src] {
		if !flag[flights[1]] {
			flag[flights[1]] = true
			tmp := findCheapestPriceHelper(flightMap, flag, flights[1], dst, price+flights[2], min, K-1)
			if tmp < min {
				min = tmp
			}
			flag[flights[1]] = false
		}
	}
	return min
}
