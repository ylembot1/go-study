package leetcode

func maxOperations(nums []int) int {
	res := 0
	t := nums[0] + nums[1]
	for i := 2; i <= len(nums)-2; {
		if nums[i]+nums[i+1] == t {
			res++
		} else {
			break
		}
		i += 2
	}
	return res + 1
}
