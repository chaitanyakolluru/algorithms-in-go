package algorithms

func LinearSearch(haystack []int, needle int) bool {
	for h := 0; h < len(haystack); h++ {
		if haystack[h] == needle {
			return true
		}
	}

	return false
}
