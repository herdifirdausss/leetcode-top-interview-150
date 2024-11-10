func twoSum(nums []int, target int) []int {
	// Space: O(n)
	s := make(map[int]int)

	// Time: O(n)
	for idx, num := range nums {
		// Time: O(1)
		if pos, ok := s[target-num]; ok {
			return []int{pos, idx}
		}
		s[num] = idx
	}
	return []int{}
}