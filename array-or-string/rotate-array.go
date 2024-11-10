// we can use this method to get O(1) space
// get step to rotate (mod n if k is more than lens num)
// rotate all (a,b,c,d,e) -> (e,d,c,b,a)
// let say rotate 2, the rotate 2 first nums
// (e,d,c,b,a) -> (d,e,c,b,a)
// rotate from k to n
// (d,e,c,b,a) -> (d,e,a,b,c)

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1) // Balikkan seluruh array
	reverse(nums, 0, k-1) // Balikkan k elemen pertama
	reverse(nums, k, n-1) // Balikkan elemen dari k hingga akhir
}

// complexity time : O(n), space : O(1)

// alternate:
// func rotate(nums []int, k int) {
//     n := len(nums)
//     k = k % n
//     rotated := append(nums[n-k:], nums[:n-k]...)
//     copy(nums, rotated)
// }

// append will O(n) times and will make new array, space O(n)
// copy will O(n) times

// complexity : O(n) + O(n) = O(n) time
// space : O(n) make new array on append
