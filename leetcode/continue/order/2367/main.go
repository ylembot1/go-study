// func arithmeticTriplets(nums []int, diff int) int {
//     res := 0

//     for i, ni := range nums {
//         j := i + 1
//         for j < len(nums) {
//             k := j + 1
//             for k < len(nums) {
//                 if nums[k] - nums[j] == diff && nums[j] - ni == diff {
//                     res += 1
//                 }
//                 k += 1
//             }
//             j += 1
//         }
//     }

//     return res
// }

func arithmeticTriplets(nums []int, diff int) (ans int) {
	set := map[int]bool{}
	for _, x := range nums {
		set[x] = true
	}

	for _, x := range nums {
		if set[x-diff] && set[x+diff] {
			ans += 1
		}
	}

	return
}