package intro

import (
	"sort"
)

func sortByHeight(a []int) []int {
	b := []int{}
	for _, value := range a {
		if value != -1 {
			b = append(b, value)
		}
	}
	sort.Ints(b)
	idx := 0
	for _, value := range b {
		for j := idx; j < len(a); j++ {
			if a[j] != -1 {
				a[j] = value
				idx = j + 1
				break
			}
		}
	}
	return a
}
