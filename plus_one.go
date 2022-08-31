package awesomeProject

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
			if i == 0 {
				var newDigits []int
				newDigits = make([]int, len(digits)+1)
				newDigits[0] = 1
				for i := 0; i < len(digits)-1; i++ {
					newDigits[i+1] = digits[i]
				}
				digits = newDigits
			}
		}
	}
	return digits
}
