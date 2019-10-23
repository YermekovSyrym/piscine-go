package quest5

func Join(s []string, t string) string {
	ans := ""
	ok := true
	for _, c := range s {
		if ok {
			ans = c
			ok = false
		} else {
			ans = ans + t + c
		}
	}
	return ans
}
