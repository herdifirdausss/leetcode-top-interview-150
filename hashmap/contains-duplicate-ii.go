func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int) // Membuat map untuk menyimpan nilai dan indeksnya
	for i, num := range nums {
		if prevIndex, found := numMap[num]; found && i-prevIndex <= k {
			return true // Jika ditemukan elemen yang sama dengan jarak indeks <= k
		}
		numMap[num] = i // Simpan indeks elemen tersebut
	}
	return false // Jika tidak ada elemen yang memenuhi syarat
}