// menggunakan dua pointer
// membuat function untuk mengecek alphanumeric
// membuat function untuk menjadi lowercase
// melakukan looping pointer, skip untuk non alphanumerik
// membandingkan karakter kiri dengan kanan, apabila tidak sama
// maka return false

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Lewati karakter non-alfanumerik di kiri
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// Lewati karakter non-alfanumerik di kanan
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// Bandingkan dengan mengabaikan case sensitivity
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}
	return true
}

func isAlphanumeric(r byte) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

func toLower(r byte) byte {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

// complexity time : O(n), space : O(1)

// alternate :

// func isPalindrome(s string) bool {
//     f := func(r rune) rune {
//         if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
//             return -1
//         }

//         return unicode.ToLower(r)
//     }
//     s = strings.Map(f, s)

//     for i := 0; i < len(s)/2; i++ {
//         if s[i] != s[len(s)-1-i] {
//             return false
//         }
//     }
//     return true
// }

// complexity time : O(n), space : O(n) because strings.Map membuat salinan data string s