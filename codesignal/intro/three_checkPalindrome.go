package intro

func checkPalindrome(inputString string) bool {
	totalLength := len(inputString) - 1
	for i := 0; i < totalLength; i++ {
		if inputString[i] != inputString[totalLength-i] {
			return false
		}
	}
	return true
}
