/*
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.07%)
 * Total Accepted:    190.5K
 * Total Submissions: 594.1K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 */

import (
	"strings"
	"unicode"
)

func lengthOfLastWord(s string) int {
	slice := strings.Split(s, " ")
	for i := len(slice) - 1; i >= 0; i-- {
		for _, runeCode := range slice[i] {
			if !unicode.IsLetter(runeCode) {
				break
			}
			return len(slice[i])
		}
	}
	return 0
}