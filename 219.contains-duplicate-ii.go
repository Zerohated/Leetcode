/*
 * [219] Contains Duplicate II
 *
 * https://leetcode.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (32.89%)
 * Total Accepted:    143.3K
 * Total Submissions: 435.8K
 * Testcase Example:  '[]\n0'
 *
 *
 * Given an array of integers and an integer k, find out whether there are two
 * distinct indices i and j in the array such that nums[i] = nums[j] and the
 * absolute difference between i and j is at most k.
 *
 */
// package main

func containsNearbyDuplicate(nums []int, k int) bool {
	for index, item := range nums {
		if index+k < len(nums) {
			for _, value := range nums[index+1 : index+k+1] {
				if value == item {
					return true
				}
			}
		} else {
			for _, value := range nums[index+1:] {
				if value == item {
					return true
				}
			}

		}
	}
	return false
}

// func main() {
// 	array := []int{}
// 	fmt.Println(containsNearbyDuplicate(array, 0))
// 	fmt.Println(containsNearbyDuplicate(array, 2))
// }
