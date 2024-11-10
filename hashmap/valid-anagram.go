
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCount := [26]int{}

	for _, c := range s {
		charCount[c-'a']++
	}

	for _, c := range t {
		if charCount[c-'a'] <= 0 {
			return false
		}
		charCount[c-'a']--
	}

	return true
}

// alternate:

// for follow up question :
// What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
// we can use this, or just replace 26 with 256

// func isAnagram(s, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	charCount := make(map[rune]int)

// 	for _, c := range s {
// 		charCount[c]++
// 	}

// 	for _, c := range t {
// 		if charCount[c] <= 0 {
// 			return false
// 		}
// 		charCount[c]--
// 	}

// 	return true
// }
