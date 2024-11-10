// iterate throguh array, we know majority is > n/2
// so we can go with this with O(1 space)
// we always ++ if same value, and -- if not same
// when 0, make new candidate
// the name is Boyer-Moore Voting Algorithm

func majorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, val := range nums {
		if count == 0 {
			candidate = val
		}
		if val == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// complexity time : O(n) space : O(1)

// alternate:
// func majorityElement(nums []int) int {
//     data := map[int]int{}
//     var majority, num int
//     for _,val:=range nums{
//         data[val]++
//         if data[val]>majority{
//             majority=data[val]
//             num=val
//         }
//     }
//     return num
// }

// with this, we have space : O(n)