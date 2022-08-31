package awesomeProject

func TwoSum(nums []int, target int) []int {
	returnArray := []int{0, 0}
	var ints map[int]int
	ints = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if myInt, ok := ints[nums[i]]; ok {
			returnArray[0] = myInt
			returnArray[1] = i
		} else {
			ints[target-nums[i]] = i
		}
	}
	return returnArray
}
