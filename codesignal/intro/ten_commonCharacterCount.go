package intro

func commonCharacterCount(s1 string, s2 string) int {
	lettersOne := make(map[rune]int)
	lettersTwo := make(map[rune]int)
	common := 0
	for _, letterOne := range s1 {
		lettersOne[letterOne] += 1
	}
	for _, letterTwo := range s2 {
		lettersTwo[letterTwo] += 1
	}
	for r, q := range lettersOne {
		q2 := lettersTwo[r]
		if q2 > 0 {
			if q > q2 {
				common += q2
				continue
			}
			common += q
		}
	}
	return common
}
