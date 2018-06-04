/*
 * [717] 1-bit and 2-bit Characters
 *
 * https://leetcode.com/problems/1-bit-and-2-bit-characters/description/
 *
 * algorithms
 * Easy (49.28%)
 * Total Accepted:    20.9K
 * Total Submissions: 42.3K
 * Testcase Example:  '[1,0,0]'
 *
 * We have two special characters. The first character can be represented by
 * one bit 0. The second character can be represented by two bits (10 or
 * 11).
 *
 * Now given a string represented by several bits. Return whether the last
 * character must be a one-bit character or not. The given string will always
 * end with a zero.
 *
 * Example 1:
 *
 * Input:
 * bits = [1, 0, 0]
 * Output: True
 * Explanation:
 * The only way to decode it is two-bit character and one-bit character. So the
 * last character is one-bit character.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * bits = [1, 1, 1, 0]
 * Output: False
 * Explanation:
 * The only way to decode it is two-bit character and two-bit character. So the
 * last character is NOT one-bit character.
 *
 *
 *
 * Note:
 * 1 .
 * bits[i] is always 0 or 1.
 *
 */
// package main

func isOneBitCharacter(bits []int) bool {
	for i := 0; i <= len(bits)-1; i++ {
		if bits[i] == 0 {
		} else if bits[i] == 1 {
			if i+1 <= len(bits)-2 {
				i++
			} else if i+1 == len(bits)-1 {
				return false
			}
		} else {
			panic("wrong value")
		}
	}
	return true
}

// func main() {
// 	bits := []int{1, 1, 1, 0}
// 	result := isOneBitCharacter(bits)
// 	fmt.Println(result)
// 	//  * Output should be: False
// 	bits = []int{1, 1, 0}
// 	result = isOneBitCharacter(bits)
// 	fmt.Println(result)
// 	//  * Output: True
// }
