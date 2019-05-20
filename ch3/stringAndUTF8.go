package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	// rune 格式就是int32，编码格式也称为UTF-32.
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		// 将当前的string按rune格式解析
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	a := []byte("afecd")
	b := []byte("adegv")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%08b", a)
}


