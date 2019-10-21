package main

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}

func main() {
	fmt.Println(RecursiveFactorial(4))
}
