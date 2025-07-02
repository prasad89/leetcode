func twoSum(nums []int, target int) []int {
    numToIndex := make(map[int]int)
    for index, num := range nums {
        complement := target - num
        if complementIndex, found := numToIndex[complement]; found {
            return []int{index, complementIndex}
        }
        numToIndex[num] = index
    }
    return []int{}
}

