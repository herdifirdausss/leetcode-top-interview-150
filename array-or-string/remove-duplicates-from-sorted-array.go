// iterate through array with two pointer, one is for the first char
// if second pointer get not same value with first pointer
// then first pointer will move to right and fill it with value second pointer
// it will remove duplicate (because it is sorted)

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

// alternate

// func removeDuplicates(nums []int) int {
//     if len(nums) <= 1 {
//         return len(nums)
//     }

//     j := 1
//     for i := 1; i < len(nums); i++ {
//         if nums[i] > nums[j-1] {
//             nums[j] = nums[i]
//             j++
//         }
//     }
//     return j
// }

// dari kode tersebut, pointer dimulai bersama pada objek ke 2
// jika objek sekarang tidak sama (lebih dari juga bisa karena sorted) dari objek sebelumnya
// pointer j sebagai penanda data yang non duplikat akan diubah menjadi nums [i]
// setelah itu pointer j akan bergerak ke kanan
// sedangkan jika nums[i]==nums[j], maka penanda data non duplikat (pointer j) tidak bergerak

// Kompleksitas Waktu
// Waktu terbaik (Best case): Jika panjang array nums adalah 0 atau 1, kita langsung mengembalikan panjang array tanpa memasuki loop, yang berarti kompleksitas waktu adalah O(1).
// Waktu rata-rata dan terburuk (Average and Worst case): Di dalam loop, kita membandingkan elemen satu per satu, yang berarti kita melakukan iterasi sebanyak n-1 kali, dengan n adalah panjang array. Untuk setiap iterasi, operasi yang dilakukan (pembandingan dan penugasan) memiliki kompleksitas O(1). Jadi, kompleksitas waktu total adalah O(n), di mana n adalah panjang array nums.

// Kompleksitas Ruang
// Ruang yang digunakan oleh fungsi ini: Fungsi ini memodifikasi array nums di tempat (in-place) dan tidak menggunakan struktur data tambahan yang besar. Hanya menggunakan variabel tambahan seperti j dan i, yang semuanya menggunakan O(1) ruang tambahan.
