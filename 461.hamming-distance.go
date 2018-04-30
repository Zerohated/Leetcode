/*
 * [461] Hamming Distance
 *
 * https://leetcode.com/problems/hamming-distance/description/
 *
 * algorithms
 * Easy (69.51%)
 * Total Accepted:    152.1K
 * Total Submissions: 218.9K
 * Testcase Example:  '1\n4'
 *
 * The Hamming distance between two integers is the number of positions at
 * which the corresponding bits are different.
 *
 * Given two integers x and y, calculate the Hamming distance.
 *
 * Note:
 * 0 ≤ x, y < 231.
 *
 *
 * Example:
 *
 * Input: x = 1, y = 4
 *
 * Output: 2
 *
 * Explanation:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 * ⁠      ↑   ↑
 *
 * The above arrows point to positions where the corresponding bits are
 * different.
 *
 *
 */
// package main

import (
	"strconv"
	"strings"
)

func hammingDistance(x int, y int) int {
	difference := x ^ y
	binaryStr := strconv.FormatInt(int64(difference), 2)
	count := strings.Count(binaryStr, "1")
	return count
}

// func main() {
// 	fmt.Println(hammingDistance(60, 13))
// }
