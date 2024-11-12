
func jump(nums []int) int {
	jumps, farthest, currentEnd := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		// Update farthest reachable position
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		// If we reach the current end, increase the jump count
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}