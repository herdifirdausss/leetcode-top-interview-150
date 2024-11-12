func threeSum(nums []int) [][]int {
	var res [][]int
	n := len(nums)

	if n < 3 {
		return res // Tidak mungkin ada triplet jika jumlah elemen kurang dari 3
	}

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// If the current number is greater than 0, break, as no valid triplet can be found
		if nums[i] > 0 {
			break
		}

		// Skip duplicate numbers for i to avoid redundant triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Two-pointer approach to find valid triplet
		j, k := i+1, n-1
		for j < k {
			total := nums[i] + nums[j] + nums[k]

			if total == 0 {
				// Found a valid triplet
				res = append(res, []int{nums[i], nums[j], nums[k]})

				// Skip duplicates for the second number
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				// Skip duplicates for the third number
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				// Move both pointers after finding a valid triplet
				j++
				k--
			} else if total < 0 {
				// Move j to the right if the sum is less than 0
				j++
			} else {
				// Move k to the left if the sum is greater than 0
				k--
			}
		}
	}

	return res
}
