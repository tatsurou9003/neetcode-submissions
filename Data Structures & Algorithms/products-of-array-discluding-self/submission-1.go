func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	left := 1
	for i:=0; i<n; i++ {
		res[i] = left
		left *= nums[i]
	}

	right := 1
	for i:=n-1; i>=0; i-- {
		res[i] *= right
		right *= nums[i]
	}
	return res
}
