package quest5

func Capitalize(s string) string {
	ok := true
	b := []rune(s)
	ln := 0
	for _, c := range b {
		if c == c {
			ln++
		}
	}
	for i := 0; i < ln; i++ {
		if i != 0 && (b[i-1] < '0' || b[i-1] > '9') && (b[i-1] < 'a' || b[i-1] > 'z') && (b[i-1] > 'Z' || b[i-1] < 'A') {
			ok = true
		}
		if ok {
			if b[i] < '0' || b[i] > '9' {
				if b[i] >= 'a' && b[i] <= 'z' {
					b[i] = rune(b[i] - 32)
				}
			}
			ok = false
		} else {
			if b[i] < '0' || b[i] > '9' {
				if b[i] >= 'A' && b[i] <= 'Z' {
					b[i] = rune(b[i] + 32)
				}
			}
		}
	}
	return string(b)
}
