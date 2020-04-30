package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // Inferência de tipos
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3

	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x

	fmt.Println(x, y)
}
