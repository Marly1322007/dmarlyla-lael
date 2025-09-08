package main

import "github.com/01-edu/z01"

func PrintMe(s string) {
	tempo := []rune(s)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(tempo[i])
	}
	z01.PrintRune('\n')
}
