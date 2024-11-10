// iterate from back, i is index, m is index for nums1, n is index for nums2
// if all m is already sorted (idx <0), then will fill remain with n value

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i, m, n := m+n-1, m-1, n-1
	for m >= 0 && n >= 0 {
		if nums2[n] >= nums1[m] {
			nums1[i] = nums2[n]
			n--
		} else {
			nums1[i] = nums1[m]
			m--
		}
		i--
	}
	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
}

// alternate:
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	i := m + n - 1
// 	m--
// 	n--
// 	for n >= 0 {
// 		if m >= 0 && nums1[m] > nums2[n] {
// 			nums1[i] = nums1[m]
// 			m--
// 		} else {
// 			nums1[i] = nums2[n]
// 			n--
// 		}
// 		i--
// 	}
// }

// Penjelasan refaktor:

// Penghapusan pengecekan if n == 0:
// Kondisi ini tidak diperlukan karena jika n bernilai 0, loop for n >= 0 tidak akan dieksekusi.

// Penggabungan kondisi m >= 0 ke dalam loop utama: Loop kedua yang hanya memproses nums2 tidak diperlukan.
// Kita bisa menanganinya dalam satu loop for n >= 0, dengan mengecek apakah m >= 0 terlebih dahulu.

// Kompleksitas Waktu (Time Complexity)
// Fungsi merge menggunakan satu loop for yang berjalan sampai semua elemen di nums2 diproses (ketika n >= 0).
// Dalam setiap iterasi, kita hanya melakukan operasi perbandingan dan pengurangan pada m atau n, yang keduanya paling banyak akan dilakukan sebanyak m + n kali.
// Oleh karena itu, kompleksitas waktu dari fungsi ini adalah
// O(m+n), di mana m adalah panjang efektif nums1 dan n adalah panjang nums2.

// Kompleksitas Ruang (Space Complexity)
// Fungsi ini hanya menggunakan beberapa variabel tambahan (seperti i, m, dan n) untuk melacak indeks dalam array.
// Tidak ada alokasi ruang tambahan yang signifikan seperti array baru atau struktur data lain.
// Sehingga kompleksitas ruang dari fungsi ini adalah
// O(1), karena hanya menggunakan ruang konstan tambahan.