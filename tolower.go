package quest5

func ToLower(s string) string {
	b := []rune(s)
	for i := 0; i < len(s); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = rune(b[i] + 32)
		}
	}
	return string(b)
}
