package leetcode

func average(salary []int) float64 {
	min, max := salary[0], salary[0]
	total := 0
	for i := 0; i < len(salary); i++ {
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
		total += salary[i]
	}

	return float64(total-min-max) / float64(len(salary)-2)
}
