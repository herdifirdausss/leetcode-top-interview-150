func isIsomorphic(s string, t string) bool {
	// Dua array untuk pemetaan karakter dari s ke t dan t ke s
	// Include all ascii (256)
	var mapS [256]int
	var mapT [256]int

	for i := 0; i < len(s); i++ {
		// Mendapatkan nilai characternya
		charS := s[i]
		charT := t[i]

		// Jika karakter belum dipetakan, petakan mereka
		if mapS[charS] == 0 && mapT[charT] == 0 {
			mapS[charS] = int(charT) + 1 // Hindari konflik dengan nilai 0
			mapT[charT] = int(charS) + 1 // Hindari konflik dengan nilai 0
		} else if mapS[charS] != int(charT)+1 || mapT[charT] != int(charS)+1 {
			return false
		}
	}

	return true
}

// alternte:

// func isIsomorphic(s string, t string) bool {
//     // Menggunakan map untuk pemetaan karakter berdasarkan nilai ASCII mereka
//     charMapS := make(map[int]int)
//     charMapT := make(map[int]int)

//     for i := 0; i < len(s); i++ {
//         charS := int(s[i]) // Nilai ASCII karakter s[i]
//         charT := int(t[i]) // Nilai ASCII karakter t[i]

//         // Pengecekan pemetaan di kedua arah sekaligus
//         if mappedS, existsS := charMapS[charS]; existsS && mappedS != charT {
//             return false
//         }
//         if mappedT, existsT := charMapT[charT]; existsT && mappedT != charS {
//             return false
//         }

//         // Membangun pemetaan karakter
//         charMapS[charS] = charT
//         charMapT[charT] = charS
//     }

//     return true
// }