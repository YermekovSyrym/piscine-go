package quest5

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(s int) {

	m := [10]int{}
	if s == 0 {
		m[s]++
	}
	for {
		if s == 0 {
			break
		}
		m[s%10]++
		s /= 10
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < m[i]; j++ {
			z01.PrintRune(rune(48 + i))
		}
	}
}
