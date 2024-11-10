func wordPattern(pattern string, s string) bool {
	// Split input string `s` into words
	words := strings.Fields(s)

	// If the number of words doesn't match the length of the pattern, return false
	if len(words) != len(pattern) {
		return false
	}

	// Create two maps for checking the bidirectional mapping
	patternToWord := make(map[rune]string)
	wordToPattern := make(map[string]rune)

	// Iterate over both the pattern and the words simultaneously
	for i, char := range pattern {
		word := words[i]

		// Check pattern -> word mapping
		if existingWord, exists := patternToWord[char]; exists && existingWord != word {
			return false
		}
		patternToWord[char] = word

		// Check word -> pattern mapping
		if existingPattern, exists := wordToPattern[word]; exists && existingPattern != char {
			return false
		}
		wordToPattern[word] = char
	}

	// If all mappings are consistent, return true
	return true
}
