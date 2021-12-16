package main

import "fmt"

// FUNC 1
// CALL BY VALUE
func add(x, y int) int {
	z := x + y
	return z
}

// CAll By reference
func swep(x, y *int) {
	a := *x
	*x = *y
	*y = a
}
func main() {
	var a, b (int)
	fmt.Scanln(&a, &b)
	fmt.Println(add(a, b))

	k := 2
	l := 3
	fmt.Println("Before Swaps : ", k, l)

	swep(&k, &l)
	fmt.Println("after Swaps : ", k, l)
}
