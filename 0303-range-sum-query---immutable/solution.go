type NumArray struct {
	prefixTable []int
}

func Constructor(nums []int) NumArray {
	prefixTable := make([]int, len(nums))
	prefixTable[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixTable[i] = prefixTable[i-1] + nums[i]
	}
	return NumArray{prefixTable: prefixTable}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixTable[right]
	} else {
		return this.prefixTable[right] - this.prefixTable[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
