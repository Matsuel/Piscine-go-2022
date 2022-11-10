package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0]
	r := []rune(arguments)
	for i := 2; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
	// si sa print toute la ligne c'est normal c'est windows
}
