/*
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (32.19%)
 * Total Accepted:    233.8K
 * Total Submissions: 726.2K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 *
 * Note:
 *
 *
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively.
 * You may assume that nums1 has enough space (size that is greater or equal to
 * m + n) to hold additional elements from nums2.
 *
 *
 * Example:
 *
 *
 * Input:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * Output: [1,2,2,3,5,6]
 *
 *
 */

// package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	n2 := nums2[:n]
	for i := m; i <= len(nums1)-1; i++ {
		nums1[i] = 0
	}
	for index, item := range n2 {
		for count := range nums1 {
			if item < nums1[count] {
				for j := len(nums1) - 1; j > count; j-- {
					nums1[j] = nums1[j-1]
				}
				nums1[count] = item
				break
			}
			if count >= m+index {
				for j := len(nums1) - 1; j > count; j-- {
					nums1[j] = nums1[j-1]
				}
				nums1[count] = item
				break
			}
		}
	}
	return
}

// func main() {
// 	nums1 := []int{1, 2, 3, 0, 0, 0}
// 	nums2 := []int{2, 5, 6}
// 	merge(nums1, 3, nums2, 3)
// 	fmt.Println(nums1)
// }
