package problems

import "strings"

/*
	345. Reverse Vowels of a String
	Description:

	Given a string s, reverse only all the vowels in the string and return it.

	The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.



	Example 1:

	Input: s = "hello"
	Output: "holle"
	Example 2:

	Input: s = "leetcode"
	Output: "leotcede"


	Constraints:

	1 <= s.length <= 3 * 105
	s consist of printable ASCII characters.
*/
///Solution:
func ReverseVowels(s string) string {
	l, r := 0, len(s)-1
	vowels := "aeiouAEIOU"
	runeSlice := []rune(s)
	for l < r {
		for l < r && !strings.Contains(vowels, string(runeSlice[l])) {
			l++
		}
		for l < r && !strings.Contains(vowels, string(runeSlice[r])) {
			r--
		}
		if l < r {
			runeSlice[l], runeSlice[r] = runeSlice[r], runeSlice[l]
			l++
			r--
		}
	}

	return string(runeSlice)
}
