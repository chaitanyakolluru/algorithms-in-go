package bigO

func sum_char_codes_n2(n string) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n); j++ {
			sum += int(n[j])
		}
	}

	return sum
}
