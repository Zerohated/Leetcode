/*
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (40.27%)
 * Total Accepted:    173.5K
 * Total Submissions: 428.9K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 *[1,5,10,10,5,1]
 * ]
 *
 *
 */
// package main

// import "fmt"

func generate(numRows int) [][]int {
	result := [][]int{}
	for i := 1; i <= numRows; i++ {
		temp := make([]int, i)
		for num := range temp {
			if num == 0 || num == len(temp)-1 {
				temp[num] = 1
			} else {
				temp[num] = result[len(result)-1][num-1] + result[len(result)-1][num]
			}
		}
		result = append(result, temp)
	}
	return result
}

// func main() {
// 	result := generate(5)
// 	for i := range result {
// 		fmt.Println(result[i])
// 	}
// }
