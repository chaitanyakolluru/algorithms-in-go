package algorithms

import "math"

func BinarySearchList(haystack []int, needle int) bool {
	var lo float64 = 0
	hi := float64(len(haystack))

	for lo < hi {
		m := math.Floor(lo + (hi-lo)/2)
		v := haystack[int(m)]

		if v == needle {
			return true
		}
		if v < needle {
			lo = m + 1
		}
		if v > needle {
			hi = m
		}
	}

	return false
}
