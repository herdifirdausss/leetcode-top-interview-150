func isIsomorphic(s string, t string) bool {
	var mapS, mapT [256]int

	for i := 0; i < len(s); i++ {
		charS, charT := s[i], t[i]

		if mapS[charS] != mapT[charT] {
			return false
		}

		mapS[charS] = i + 1
		mapT[charT] = i + 1
	}
	return true
}

// alternte:

// func isIsomorphic(s string, t string) bool {
//     mapS := make(map[rune]rune)
//     mapT := make(map[rune]rune)

//     for i, charS := range s {
//         charT := rune(t[i])

//         if mapped, ok := mapS[charS]; ok && mapped != charT {
//             return false
//         }
//         if mapped, ok := mapT[charT]; ok && mapped != charS {
//             return false
//         }

//         mapS[charS] = charT
//         mapT[charT] = charS
//     }
//     return true
// }
