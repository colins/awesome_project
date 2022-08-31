package awesomeProject

func Rotate(nums []int, k int) {
	if k == 0 || len(nums) == 1 {
		return
	}
	if k > len(nums) {
		k = k % len(nums)
	}
	offset := len(nums) - k
	var numsCopy []int
	numsCopy = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numsCopy[i] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = numsCopy[offset]
		if offset == len(nums)-1 {
			offset = 0
		} else {
			offset++
		}
	}
}
