package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	ok := 0
	for _, c := range arg {
		cur := 0
		for _, v := range c {
			cur = cur*10 + int(v-'0')
		}
		if c == "--upper" {
			ok = 1
		}
		if c == "--upper" || cur > 26 {
			z01.PrintRune(' ')
			continue
		}
		z01.PrintRune(rune((cur + 'a' - 1 - ok*32)))
	}
	z01.PrintRune('\n')
}
