package intro

import "fmt"

func isLucky(n int) bool {
	number := fmt.Sprint(n)
	length := len(number)
	half := length / 2
	sumOne := 0
	sumTwo := 0
	for i := 0; i < half; i++ {
		sumOne += int(number[i])
	}
	for i := half; i < length; i++ {
		sumTwo += int(number[i])
	}

	return sumOne == sumTwo
}
