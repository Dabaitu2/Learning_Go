package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func pointerReverse(s *[5]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func removeRepeat(s []string) []string {
	var i int
	var left string
	for _, v := range s {
		if left == string(v) {
			continue
		}
		s[i] = v
		left = string(v)
		i++
	}
	return s[:i]
}

func cutUnicode(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		index := i + 1
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			copy(s[i:], s[index:])
		}
	}
	return s
}

// 不重新分配内存实现反转utf-8字符串
func utf8Reverse(s []byte) []byte {
	for i := len(s); i > 0; {
		// 读取当前的rune字符
		r, size := utf8.DecodeRuneInString(string(s[0:]))
		fmt.Println(string(r))
		fmt.Println(size)
		// 把除了最后一个rune外全部前移，顶出第一个参数
		copy(s[0:i], s[0+size:i])
		// 把获得的临时变量赋值给最后的参数，完成放置
		copy(s[i-size:i], []byte(string(r)))
		// 缩小交换范围
		i -= size
	}
	fmt.Println(string(s))
	return s
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	pointerReverse(&a)
	fmt.Println(a)
	b := [...]string{"a", "b", "b", "a", "c", "b"}
	fmt.Println(removeRepeat(b[:]))
	c := "  asda  asda"
	fmt.Printf("%s\n", cutUnicode([]byte(c)))
	d := "阿大大的asda"
	utf8Reverse([]byte(d))
}
