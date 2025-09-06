func findIntersectionValues(nums1 []int, nums2 []int) []int {
	set1 := map[int]struct{}{}
	for _, n := range nums1 {
		set1[n] = struct{}{}
	}

	set2 := map[int]struct{}{}
	for _, n := range nums2 {
		set2[n] = struct{}{}
	}

	ans := make([]int, 2)

	for _, n := range nums1 {
		if _, ok := set2[n]; ok {
			ans[0] += 1
		}
	}

	for _, n := range nums2 {
		if _, ok := set1[n]; ok {
			ans[1] += 1
		}
	}

	return ans
}