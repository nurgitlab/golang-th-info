package sum

func SumFloat(args ...float64) float64 {
	ans := float64(0)
	for _, arg := range args {
		ans += arg
	}

	return ans
}
