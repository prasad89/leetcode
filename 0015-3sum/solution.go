func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			currSum := nums[i] + nums[j] + nums[k]
			if currSum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if currSum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return res
}
