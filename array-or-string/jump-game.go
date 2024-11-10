// iterate through array from the back
// if the index + value of the index can reach the goal,
// so the index can be goal too, because it will make us reach goal
// it will iterate back, if goal is on index 0
// it mean from index 0, we can reach the goal

// another solution : https://leetcode.com/problems/jump-game/discuss/3878729/go-multiple-solutions-in-go-golang-100

func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0
}

// complexity time = O(n), space = O(1)

// alternate:

// func canJump(nums []int) bool {
//     for pos, farthest := 0, 0; pos < len(nums); pos++ {
//         if pos > farthest {
//             return false
//         }
//         farthest = max(farthest, pos + nums[pos])
//     }
//     return true
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// with this code, we iterate position from the front
// if position + value on the position can > curr farthest
// then change the value
// if pos > farthest, it mean not possible to be here
// so, if pos>farthest, and looping not finished
// it mean we can't reach the end, because end of pos is end of index (goal)