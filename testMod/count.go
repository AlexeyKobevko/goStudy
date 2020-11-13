package count

func Count(s string, r rune) int {
	var cnt int
	for _, i := range s {
		if i == r {
			cnt += 1
		}
	}
	return cnt
}
