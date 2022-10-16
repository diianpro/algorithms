package sumOddLengthSubarrays

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, bank := range accounts {
		sum := 0
		for _, bill := range bank {
			sum += bill
		}
		if sum >= max {
			max = sum
		}
	}
	return max
}
