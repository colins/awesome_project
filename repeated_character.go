package awesomeProject

func RepeatedCharacter(s string) byte {
	var charMap [26]bool
	for _, char := range s {
		if charMap[char-'a'] {
			return byte(char)
		}
		charMap[char-'a'] = true
	}
	return 'a'
}
