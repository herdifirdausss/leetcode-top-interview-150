// looping through array, will fill nums index k when num is not val
// it will always assign even if there's not val on the array

func removeElement(nums []int, val int) int {
	var k int
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k

}

// alternate:
// func removeElement(nums []int, val int) int {
//     left, right := 0, len(nums)-1
//     for left <= right {
//         if nums[left] == val {
//             nums[left] = nums[right]
//             right--
//         } else {
//             left++
//         }
//     }
//     return left
// }

// if there's not have val on the array, it will just loop
// not assign everything not val
// run time will be fast because it not assign sometime

// but it will have different output, if output matter, choose wisely

// Kompleksitas Waktu: Kedua implementasi memiliki kompleksitas waktu O(n) pada best dan worst case, namun implementasi kedua dapat lebih efisien dalam runtime jika elemen val jarang muncul.
// Kompleksitas Ruang: Keduanya O(1), tanpa alokasi tambahan untuk array.