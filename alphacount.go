package piscine

func AlphaCount(str string) int {
	s := 0
	for _, i := range str {
		if (i >= 'a' && i <= 'z') || (i <= 'Z' && i >= 'A') {
			s++
		}
	}
	return s
}
