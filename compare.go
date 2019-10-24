package quest5

func Compare(s, t string) int {
	ln := 0
	ln2 := 0
	for _, c := range s {
		if c == c {
			ln++
		}
	}

	for _, c := range t {
		if c == c {
			ln2++
		}
	}
	x := ln
	if ln > ln2 {
		x = ln2
	}
	for i := 0; i < x; i++ {
		if s[i] > t[i] {
			return 1
		}
		if s[i] < t[i] {
			return -1
		}
	}
	if ln > ln2 {
		return 1
	}
	if ln < ln2 {
		return -1
	}
	return 0
}
