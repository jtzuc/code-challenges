package intro

func centuryFromYear(year int) int {
	remaining := year % 100
	if remaining == 0 {
		return year / 100
	}
	return (year / 100) + 1

}
