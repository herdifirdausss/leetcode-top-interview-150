func groupAnagrams(words []string) [][]string {
	cache := make(map[[26]int]int)
	result := make([][]string, 0)
	for _, s := range words {
		charCount := [26]int{}

		for _, c := range s {
			charCount[c-'a']++
		}
		if idx, ok := cache[charCount]; ok {
			// we append the word to the right slice in the result
			result[idx] = append(result[idx], words[i])
		} else {
			result = append(result, []string{words[i]})
			cache[charCount] = len(result) - 1
		}
	}
	return result
}