package sep152024

func minValidStrings(words []string, target string) int {
	n := len(target)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	prefixes := make(map[string]bool)
	for _, word := range words {
		for i := 1; i <= len(word); i++ {
			prefixes[word[:i]] = true
		}
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] != -1 && prefixes[target[j:i]] {
				if dp[i] == -1 || dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	return dp[n]
}
