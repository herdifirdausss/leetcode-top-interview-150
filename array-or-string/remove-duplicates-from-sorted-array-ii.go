// dari kode tersebut, pointer dimulai bersama pada objek ke 3
// dikarenakan jika objek pertama dan kedua sama pun, biarkan saja atau jika hanya ada 2 objek maka return saja
// jika objek sekarang tidak sama (lebih dari juga bisa karena sorted) dari objek kedua sebelumnya
// pointer j sebagai penanda data yang non duplikat akan diubah menjadi nums [i]
// kenapa dua objek sebelumnya? jika objek sebelumnya sama dengan objek sekarang, otomatis, ada 3 duplikat dalam hal ini
// setelah itu pointer j akan bergerak ke kanan
// sedangkan jika nums[i]==nums[j], maka penanda data non duplikat (pointer j) tidak bergerak

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

// Kompleksitas waktu : O(n) Kompleksitas space : O(1)
