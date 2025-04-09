package main

import (
	"sort"
	"strconv"
)

func stableMountains(height []int, threshold int) []int {
	res := []int{}
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			res = append(res, i)
		}
	}
	return res
}

func minElement(nums []int) int {
	res := []int{}
	for _, val := range nums {
		num := 0
		for _, part := range strconv.Itoa(val) {
			numC, _ := strconv.Atoi(string(part))
			num += numC
		}
		res = append(res, num)
	}
	min := res[0]
	for _, val := range res {
		if val < min {
			min = val
		}
	}
	return min
}

func maximumTotalSum(maximumHeight []int) int64 {
	sort.Ints(maximumHeight)

	var totalSum int64 = 0
	currHeight := maximumHeight[len(maximumHeight)-1]

	for i := len(maximumHeight) - 1; i >= 0; i-- {
		if maximumHeight[i] < currHeight {
			currHeight = maximumHeight[i]
		}

		if currHeight <= 0 {
			return -1
		}

		totalSum += int64(currHeight)

		currHeight--
	}

	return totalSum
}

func minElement2(nums []int) int {
	minSum := int(^uint(0) >> 1) // Set initial minSum to the largest possible int

	// Helper function to compute sum of digits
	sumDigits := func(num int) int {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		return sum
	}

	// Iterate over nums and find the minimum sum of digits
	for _, val := range nums {
		digitSum := sumDigits(val)
		if digitSum < minSum {
			minSum = digitSum
		}
	}

	return minSum
}

func validSequence(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)
	tenvoraliq := []int{} // stores the indices of the valid sequence

	i, j := 0, 0
	modified := false // to track if we've already modified one character

	// Iterate through word1 while trying to form word2
	for i < n && j < m {
		if word1[i] == word2[j] {
			// Direct match, no need to modify
			tenvoraliq = append(tenvoraliq, i)
			j++
		} else if !modified {
			// Modify one character to satisfy the "almost equal" condition
			tenvoraliq = append(tenvoraliq, i)
			modified = true
			j++
		}
		i++
	}

	// If we successfully matched all characters from word2, return the result
	if j == m {
		return tenvoraliq
	}

	// Otherwise, return an empty array as no valid sequence exists
	return []int{}
}

func main() {
}
