func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		currSum := numbers[l] + numbers[r]
		if currSum == target {
			return []int{l + 1, r + 1}
		}
		if currSum < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
