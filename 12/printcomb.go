package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				for k := j + 1; k <= '9'; k++ {
					if i < c {
						z01.PrintRune('j')
					} else if b < d {
						z01.PrintRune('c')
					}
					if i != '7' || j != '8' || k != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}