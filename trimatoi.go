package quest5

func check(s string) bool {

	ok := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			ok = true
			break
		}
	}
	return ok
}
func TrimAtoi(str string) int {
	x := 0

	change := false

	for _, c := range str {
		if c >= '0' && c <= '9' {
			break
		}
		if c == '-' {
			change = true
		}
		if c == '+' {
			change = false
		}
	}
	if check(str) == true {
		for _, c := range str {

			cnt := 0
			if c >= '0' && c <= '9' {
				for i := '1'; i <= c; i++ {
					cnt++
				}
				x = x*10 + cnt
			}
		}
	}
	if change == true {
		x *= -1
	}
	return x
}
