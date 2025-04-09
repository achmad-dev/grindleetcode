package sep152024

func minValidStrings2(words []string, target string) int {
	n := len(target)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
	}

	validPrefixes := make(map[string]bool)
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			validPrefixes[word[:i]] = true
		}
	}

	for i := 1; i <= n; i++ {
		for j := i; j > 0; j-- {
			if validPrefixes[target[j-1:i]] {
				if j == 1 {
					dp[i] = 1
				} else if dp[j-1] != n+1 {
					dp[i] = min(dp[i], dp[j-1]+1)
				}
			}
		}
	}

	if dp[n] == n+1 {
		return -1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
