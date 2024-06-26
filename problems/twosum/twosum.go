package twosum

func BruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func Optimal(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for i, val := range nums {
		diff := target - val
		if idx, found := hashMap[diff]; found {
			return []int{idx, i}
		}
		hashMap[val] = i
	}
	return []int{}
}
