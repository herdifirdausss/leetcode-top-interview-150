func maxArea(height []int) int {
	i, j := 0, len(height)-1
	currMax := 0
	for i <= j {
		currMax = max(currMax, (j-i)*min(height[i], height[j]))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return currMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}