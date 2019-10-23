package quest5

func IsAlpha(s string) bool {
	for _, c := range s {
		if (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c > 'Z' || c < 'A') {
			return false
		}
	}
	return true
}
