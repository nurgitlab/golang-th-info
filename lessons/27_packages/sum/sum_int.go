package sum

func SumInt(nums ...int) int {
	ans := int(0)
	for _, num := range nums {
		ans += num * 2
	}

	return ans
}
