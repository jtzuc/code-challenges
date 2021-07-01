package intro

func allLongestStrings(inputArray []string) []string {
	lengths := make(map[int][]string)
	maxLength := 0
	for _, str := range inputArray {
		length := len(str)
		lengths[length] = append(lengths[length], str)
		if length > maxLength {
			maxLength = length
		}
	}
	return lengths[maxLength]
}
