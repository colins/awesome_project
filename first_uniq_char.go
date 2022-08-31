package awesomeProject

func FirstUniqChar(s string) int {
	var charMap [26]int
	for _, char := range s {
		charMap[char-'a']++
	}
	for i, char := range s {
		if charMap[char-'a'] == 1 {
			return i
		}
	}
	return -1
}
