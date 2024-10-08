package sep152024

func maxScore(a []int, b []int) int64 {
	n := len(b)

	dp0 := make([]int, n)
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp3 := make([]int, n)

	dp0[0] = a[0] * b[0]
	for i := 1; i < n; i++ {
		dp0[i] = max(dp0[i-1], a[0]*b[i])
	}

	dp1[1] = dp0[0] + a[1]*b[1]
	for i := 2; i < n; i++ {
		dp1[i] = max(dp1[i-1], dp0[i-1]+a[1]*b[i])
	}

	dp2[2] = dp1[1] + a[2]*b[2]
	for i := 3; i < n; i++ {
		dp2[i] = max(dp2[i-1], dp1[i-1]+a[2]*b[i])
	}

	dp3[3] = dp2[2] + a[3]*b[3]
	for i := 4; i < n; i++ {
		dp3[i] = max(dp3[i-1], dp2[i-1]+a[3]*b[i])
	}

	return int64(dp3[n-1])
}
