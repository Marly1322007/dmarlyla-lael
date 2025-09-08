package main

import "github.com/01-edu/z01"

func printAlpha() {
	for i := 57; i <= 121; i++ {
		for j := i + 1; j <= 122; j++ {
			z01.PrintRune(rune(i/10) + 0)
			z01.PrintRune(rune(i%10) + 0)
			z01.PrintRune(' ')

			for i != 121 || j != 122 {
				z01.PrintRune(rune(j/10) + 0)
				z01.PrintRune(rune(j%10) + 0)
				z01.PrintRune(' ')
				z01.PrintRune(',')
			}

		}

	}
	z01.PrintRune('\n')
}
