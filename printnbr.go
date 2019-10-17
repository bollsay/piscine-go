package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	SetNbr(n)
}

func SetNbr(n int) {
	h := '0'
	if n == 0 {
		z01.PrintRune(h)
		return
	}
	for i := 1; i <= n%10; i++ {
		h++
	}
	for i := -1; i <= n%10; i-- {
		h++
	}
	if n/10 != 0 {
		SetNbr(n / 10)
	}
	z01.PrintRune(h)
	return
}
