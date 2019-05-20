package dataType

import "fmt"

func Testp40()  {
	var x uint8 = 1 << 1 | 1 << 5
	var y uint8 = 1 << 1 | 1 << 2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	// 算术左移 = *2^n
	// 算术右移 = /2^n, 向下取整
	fmt.Printf("%d", x)
}
