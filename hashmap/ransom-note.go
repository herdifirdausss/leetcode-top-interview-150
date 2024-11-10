func canConstruct(ransomNote string, magazine string) bool {
	charCount := [26]int{}

	// Mengisi jumlah karakter dari `magazine`
	for _, c := range magazine {
		charCount[c-'a']++
	}

	// Mengecek apakah `ransomNote` bisa dibentuk dari `magazine`
	for _, c := range ransomNote {
		if charCount[c-'a'] <= 0 {
			return false
		}
		charCount[c-'a']--
	}

	return true
}

// complexity time: O(n+m) space: O(1) (26)

// alternate:

// func canConstruct(ransomNote string, magazine string) bool {
//     magaHash := make(map[rune]int)

//     // Menghitung jumlah setiap karakter dalam `magazine`
//     for _, c := range magazine {
//         magaHash[c]++
//     }

//     // Memeriksa setiap karakter dalam `ransomNote`
//     for _, c := range ransomNote {
//         if magaHash[c] <= 0 {
//             return false
//         }
//         magaHash[c]--
//     }

//     return true
// }

// complexity time: O(n+m) space: O(k) k<=26