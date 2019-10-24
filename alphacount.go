package piscine

func AlphaCount(s string) int {
	cnt := 0
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c <= 'Z' && c >= 'A') {
			cnt++
		}
	}
	return cnt
}
