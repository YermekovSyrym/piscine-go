package quest5

func LastRune(s string) rune {
	var ans rune
	ln := 0
	for _, c := range s {
		if c == c {
			ln++
		}
	}
	cnt := 0
	for _, c := range s {
		if cnt == ln-1 {
			return c
		}
		cnt++
	}
	return ans
}
