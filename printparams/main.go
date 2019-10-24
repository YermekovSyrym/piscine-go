package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := os.Args
	for _, i := range i[1:] {
		for _, a := range i {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
