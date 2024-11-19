package bigO

func sum_char_codes(n string) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += int(n[i])
	}

	return sum
}
