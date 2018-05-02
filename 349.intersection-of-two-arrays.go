/*
 * [349] Intersection of Two Arrays
 *
 * https://leetcode.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (48.30%)
 * Total Accepted:    131.1K
 * Total Submissions: 271.4K
 * Testcase Example:  '[]\n[]'
 *
 *
 * Given two arrays, write a function to compute their intersection.
 *
 *
 * Example:
 * Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
 *
 *
 * Note:
 *
 * Each element in the result must be unique.
 * The result can be in any order.
 *
 *
 */

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	mapN1, mapN2 := make(map[int]bool), make(map[int]bool)
	for _, item := range nums1 {
		mapN1[item] = true
	}
	for _, item := range nums2 {
		mapN2[item] = true
	}
	for num := range mapN1 {
		if mapN2[num] == true {
			result = append(result, num)
		}
	}

	return result
}
