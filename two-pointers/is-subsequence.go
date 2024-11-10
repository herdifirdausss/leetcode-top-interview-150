// loop through array, if same, then check next char
// if pointing j (counting same char) same with length s, then true
// else false

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
		}
		if j == len(s) {
			return true
		}
	}
	return false

}

// complexity : O(n), O(1)

// alternate:

// func isSubsequence(s string, t string) bool {
//     i, j := 0, 0
//     for i < len(s) && j < len(t) {
//         if s[i] == t[j] {
//             i++
//         }
//         j++
//     }
//     return i == len(s)
// }

// Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109,
// and you want to check one by one to see if t has its subsequence.
// In this scenario, how would you change your code?

// Preprocess function to create a map of character positions in t
// func preprocess(t string) map[rune][]int {
//     posMap := make(map[rune][]int)
//     for i, c := range t {
//         posMap[c] = append(posMap[c], i)
//     }
//     return posMap
// }

// // Binary search helper function
// func binarySearch(arr []int, target int) int {
//     left, right := 0, len(arr)-1
//     for left <= right {
//         mid := left + (right-left)/2
//         if arr[mid] > target {
//             right = mid - 1
//         } else {
//             left = mid + 1
//         }
//     }
//     return left
// }

// // Function to check if s is a subsequence of t using preprocessed data
// func isSubsequence(s string, t string) bool {
//     posMap := preprocess(t)
//     currPos := -1
//     for _, c := range s {
//         if posList, ok := posMap[c]; ok {
//             // Find the smallest position in posList that is greater than currPos
//             idx := binarySearch(posList, currPos)
//             if idx == len(posList) {
//                 return false
//             }
//             currPos = posList[idx]
//         } else {
//             return false
//         }
//     }
//     return true
// }

// Preprocessing (preprocess function): Kita membuat peta posMap yang menyimpan daftar posisi masing-masing karakter dalam string t. Misalnya,
// untuk t = "abcde", peta ini akan berisi:
// {
//     'a': [0],
//     'b': [1],
//     'c': [2],
//     'd': [3],
//     'e': [4]
// }

// Binary Search: Ketika memeriksa apakah s merupakan subsequence dari t, kita mencari setiap karakter dalam s menggunakan pencarian biner untuk menemukan posisi yang lebih besar dari posisi karakter sebelumnya di t. Fungsi binarySearch memastikan pencarian posisi karakter di t dilakukan secara efisien dalam waktu logaritmik.