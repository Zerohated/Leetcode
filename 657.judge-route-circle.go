/*
 * [657] Judge Route Circle
 *
 * https://leetcode.com/problems/judge-route-circle/description/
 *
 * algorithms
 * Easy (68.53%)
 * Total Accepted:    75.4K
 * Total Submissions: 110K
 * Testcase Example:  '"UD"'
 *
 *
 * Initially, there is a Robot at position (0, 0). Given a sequence of its
 * moves, judge if this robot makes a circle, which means it moves back to the
 * original place.
 *
 *
 *
 * The move sequence is represented by a string. And each move is represent by
 * a character. The valid robot moves are R (Right), L (Left), U (Up) and D
 * (down). The output should be true or false representing whether the robot
 * makes a circle.
 *
 *
 * Example 1:
 *
 * Input: "UD"
 * Output: true
 *
 *
 *
 * Example 2:
 *
 * Input: "LL"
 * Output: false
 *
 *
 */
func judgeCircle(moves string) bool {
	position := []int{0, 0}
	for _, value := range moves {
		switch value {
		case 'U':
			position[1] += 1
		case 'D':
			position[1] -= 1
		case 'L':
			position[0] -= 1
		case 'R':
			position[0] += 1
		}
	}
	if position[0] == 0 && position[1] == 0 {
		return true
	}
	return false
}
