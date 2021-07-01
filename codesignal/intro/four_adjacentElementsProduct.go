package intro

func adjacentElementsProduct(inputArray []int) int {
	maxProduct := 0
	maxIdx := len(inputArray)
	for i := 0; i < maxIdx-1; i++ {
		product := inputArray[i] * inputArray[i+1]
		if i == 0 {
			maxProduct = product
			continue
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct
}
