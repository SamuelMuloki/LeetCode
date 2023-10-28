package solutions

type NumArray struct {
	nums []int
}

func SumRangeConstructor(nums []int) NumArray {
	arr := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		arr[i+1] = arr[i] + nums[i]
	}

	return NumArray{arr}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
