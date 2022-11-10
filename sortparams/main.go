package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	tab := []string{}
	for i := 1; i < len(arguments); i++ {
		tab = append(tab, arguments[i])
	}
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab); j++ {
			tmp := tab[i]
			if tab[i] < tab[j] {
				tab[i] = tab[j]
				tab[j] = tmp
			}
		}
	}
	for i := 0; i < len(tab); i++ {
		for _, ch := range tab[i] {
			z01.PrintRune(rune(ch))
		}
		z01.PrintRune('\n')
	}
}
