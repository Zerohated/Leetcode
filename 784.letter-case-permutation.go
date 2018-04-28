/*
 * [800] Letter Case Permutation
 *
 * https://leetcode.com/problems/letter-case-permutation/description/
 *
 * algorithms
 * Easy (52.30%)
 * Total Accepted:    10.7K
 * Total Submissions: 20.5K
 * Testcase Example:  '"a1b2"'
 *
 * Given a string S, we can transform every letter individually to be lowercase
 * or uppercase to create another string.  Return a list of all possible
 * strings we could create.
 *
 *
 * Examples:
 * Input: S = "a1b2"
 * Output: ["a1b2", "a1B2", "A1b2", "A1B2"]
 *
 * Input: S = "3z4"
 * Output: ["3z4", "3Z4"]
 *
 * Input: S = "12345"
 * Output: ["12345"]
 *
 *
 * Note:
 *
 *
 * S will be a string with length at most 12.
 * S will consist only of letters or digits.
 *
 */

import (
	"bytes"
	"strings"
	"unicode"
)

func letterCasePermutation(S string) []string {
	if S == "" {
		return []string{S}
	}
	rest := letterCasePermutation(S[:len(S)-1])
	first := []byte{S[len(S)-1]}
	if unicode.IsLetter(bytes.Runes(first)[0]) {
		tempSlice := []string{}
		for _, c := range rest {
			tempSlice = append(tempSlice, c+strings.ToLower(string(S[len(S)-1])), c+strings.ToUpper(string(S[len(S)-1])))
		}
		return tempSlice
	}
	tempSlice := []string{}
	for _, c := range rest {
		tempSlice = append(tempSlice, c+string(S[len(S)-1]))
	}
	return tempSlice
}
