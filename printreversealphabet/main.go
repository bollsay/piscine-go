package main

import "github.com/01-edu/z01"

func main() {
	for i := rune(122); i >= rune(97); i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune(rune(10))
}
