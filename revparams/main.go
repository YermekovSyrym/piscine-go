package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	i := 0
	for a := range argument {
		i = a
	}
	for b := i; b >= 1; b-- {
		for _, a := range argument[b] {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
