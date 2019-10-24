package quest5

func ToUpper(s string) string {
	b := []rune(s)
	for i := 0; i < len(s); i++ {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = rune(b[i] - 32)
		}
	}
	return string(b)
}
