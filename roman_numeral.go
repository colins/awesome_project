package awesomeProject

func RomanToInt(s string) int {
	var numerals map[string]int
	numerals = make(map[string]int)
	numerals["I"] = 1
	numerals["V"] = 5
	numerals["X"] = 10
	numerals["L"] = 50
	numerals["C"] = 100
	numerals["D"] = 500
	numerals["M"] = 1000

	chars := []rune(s)
	total := 0
	previousChar := ""
	for i := len(chars) - 1; i >= 0; i-- {
		currentNum := numerals[string(chars[i])]
		if currentNum < numerals[previousChar] {
			total -= currentNum
		} else {
			total += currentNum
		}
		previousChar = string(chars[i])
	}
	return total
}
