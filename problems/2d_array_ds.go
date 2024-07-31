package problems

import "math"

// from hackerrank

/*
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func HourglassSum(arr [][]int32) int32 {
	// Write your code here
	max := int32(math.MinInt32)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			//a b c
			//  d
			//e f g
			// top because we need 3 element so we slice to 3
			top := sumArr(arr[i][j : j+3])
			//for the middle we need only the middle value
			mid := arr[i+1][j+1]
			//bottom
			btm := sumArr(arr[i+2][j : j+3])
			sum := top + mid + btm
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func sumArr(arr []int32) int32 {
	var sum int32
	for _, val := range arr {
		sum += val
	}
	return sum
}
