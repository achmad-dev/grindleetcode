package problems

/*
	There is a large pile of socks that must be paired by color. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

	Example


	There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .

	Function Description

	Complete the sockMerchant function in the editor below.

	sockMerchant has the following parameter(s):

	int n: the number of socks in the pile
	int ar[n]: the colors of each sock
	Returns

	int: the number of pairs
	Input Format

	The first line contains an integer , the number of socks represented in .
	The second line contains  space-separated integers, , the colors of the socks in the pile.

	Constraints

	where
	Sample Input

	STDIN                       Function
	-----                       --------
	9                           n = 9
	10 20 20 10 10 30 50 10 20  ar = [10, 20, 20, 10, 10, 30, 50, 10, 20]
	Sample Output

	3
	Explanation


	There are three pairs of socks.
*/

import "sort"

// / solution
func SockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	if len(ar) == 0 {
		return 0
	}
	var res int32 = 0
	sort.Slice(ar, func(i, j int) bool { return ar[i] < ar[j] })
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] == ar[i+1] {
			res++
			//skip next because we already count it
			i++
		}
	}
	return res
}
