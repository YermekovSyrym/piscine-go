package quest5

func IsPrintable(s string) bool {
	for _, c := range s {
		if c < 32 {
			return false
		}
	}
	return true
}
