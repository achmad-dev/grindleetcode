package movezero

// move zero to the end of array without changing the order of non zero elements
// input [1,0,2,3,0,-4]
// output [1,2,3,-4,0,0]

func MoveZeroBruteForce(nums []int) []int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l += 1
		}
	}
	return nums
}
