func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Sort array
	sort.Ints(nums)
	maxLen := 1
	currentStreak := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			// Skip duplicate elements
			continue
		} else if nums[i] == nums[i-1]+1 {
			// Consecutive number
			currentStreak++
		} else {
			// Reset streak
			maxLen = max(maxLen, currentStreak)
			currentStreak = 1
		}
	}

	// Handle the last streak
	return max(maxLen, currentStreak)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// alternate:

// func longestConsecutive(nums []int) int {
//     numSet := make(map[int]struct{}) // Set untuk menyimpan elemen unik
//     maxLen := 0

//     // Masukkan semua elemen ke dalam set (map untuk kecepatan akses)
//     for _, num := range nums {
//         numSet[num] = struct{}{}
//     }

//     for num := range numSet {
//         // Jika num-1 tidak ada dalam set, itu adalah awal urutan berturut-turut
//         if _, exists := numSet[num-1]; !exists {
//             currentNum := num
//             currentStreak := 1

//             // Cari elemen berturut-turut berikutnya
//             for {
//                 nextNum := currentNum + 1
//                 if _, exists := numSet[nextNum]; exists {
//                     currentNum = nextNum
//                     currentStreak++
//                 } else {
//                     break
//                 }
//             }

//             // Memperbarui panjang urutan berturut-turut maksimum
//             if currentStreak > maxLen {
//                 maxLen = currentStreak
//             }
//         }
//     }

//     return maxLen
// }
