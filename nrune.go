package quest5

func NRune(s string, n int) rune {
	cnt := 0
	var ans rune
	for _, c := range s {
		if cnt == n-1 {
			return c
		}
		cnt++
	}
	return ans
}
