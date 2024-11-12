func isHappy(n int) bool {
	seen := make(map[int]bool) // Hash Map untuk melacak angka yang sudah muncul
	for n != 1 {
		n = sumOfSquareOfDigits(n)
		if seen[n] { // Jika n sudah pernah muncul, kita berada dalam siklus
			return false
		}
		seen[n] = true // Tandai n sebagai sudah muncul
	}
	return true // Jika n mencapai 1, maka ini adalah happy number
}

func sumOfSquareOfDigits(n int) int {
	var sum int = 0
	for n > 0 {
		dig := n % 10
		sum += dig * dig
		n /= 10
	}
	return sum
}

// Space : O(n)

// alternate:

// func isHappy(n int) bool {
//     slow, fast := n, sumOfSquareOfDigits(n)
//     for n != 1 && slow != fast {
//         slow, fast = sumOfSquareOfDigits(slow), sumOfSquareOfDigits(sumOfSquareOfDigits(fast))
//     }
//     return fast == 1
// }

// func sumOfSquareOfDigits(n int) int {
//     var sum int = 0
//     for n > 0 {
//         dig := n % 10
//         sum += dig * dig
//         n /= 10
//     }
//     return sum
// }

// Space : O(1)