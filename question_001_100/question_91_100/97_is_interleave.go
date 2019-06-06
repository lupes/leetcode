package question_91_100

func isInterleave1(s1 string, s2 string, s3 string) bool {
	start1, start2, start3, size1, size2, size3 := 0, 0, 0, len(s1), len(s2), len(s3)
	if size1+size2 != size3 {
		return false
	} else {
		var flags = make([]int32, size3)
		for start1+start2 <= size3 {
			if start1+start2 == size3 {
				return true
			}
			if flags[start3] == 0 {
				flags[start3] = 1
				if start1 < size1 && s1[start1] == s3[start3] {
					start1 += 1
					start3 += 1
					continue
				}
			}

			if flags[start3] == 1 {
				flags[start3] = 2
				if start2 < size2 && s2[start2] == s3[start3] {
					start2 += 1
					start3 += 1
					continue
				}
			}
			if start1 == 0 && start2 == 0 && start3 == 0 {
				return false
			}
			if flags[start3-1] == 1 {
				start1 -= 1
			} else {
				start2 -= 1
			}
			start3 -= 1
			flags[start3+1] = 0
		}
		return false
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[m][n]
}
