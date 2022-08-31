package awesomeProject

import "sort"

func IntToRoman(num int) string {
	var numerals map[int]string
	numerals = make(map[int]string)
	numerals[1000] = "M"
	numerals[900] = "CM"
	numerals[500] = "D"
	numerals[400] = "CD"
	numerals[100] = "C"
	numerals[90] = "XC"
	numerals[50] = "L"
	numerals[40] = "XL"
	numerals[10] = "X"
	numerals[9] = "IX"
	numerals[5] = "V"
	numerals[4] = "IV"
	numerals[1] = "I"
	romanResult := ""
	var keys []int
	for key := range numerals {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for i := len(keys) - 1; i >= 0; i-- {
		for num >= keys[i] {
			romanResult += numerals[keys[i]]
			num -= keys[i]
		}

	}
	//for _, intValue := range keys {
	//	for num >= intValue {
	//		romanResult += numerals[intValue]
	//		num -= intValue
	//	}
	//}

	return romanResult
}
