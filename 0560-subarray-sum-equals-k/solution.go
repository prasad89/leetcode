func subarraySum(nums []int, k int) int {
	prefixMap := make(map[int]int)
	sum, count := 0, 0
	prefixMap[0] = 1
	for _, num := range nums {
		sum += num
		if freq, found := prefixMap[sum-k]; found {
			count += freq
		}
		prefixMap[sum]++
	}
	return count
}
