package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

/*
	%v	值的默认格式。当打印结构体时，“加号”标记（%+v）会添加字段名
	%#v　相应值的Go语法表示
	%T	相应值的类型的Go语法表示
	%%	字面上的百分号，并非值的占位符　

布尔值：
%t	 true 或 false
整数值：

%b	二进制表示
%c	相应Unicode码点所表示的字符
%d	十进制表示
%o	八进制表示
%q	单引号围绕的字符字面值，由Go语法安全地转义
%x	十六进制表示，字母形式为小写 a-f
%X	十六进制表示，字母形式为大写 A-F
%U	Unicode格式：U+1234，等同于 "U+%04X"
浮点数及复数：

%b	无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat中的 'b' 转换格式一致。例如 -123456p-78
%e	科学计数法，例如 -1234.456e+78
%E	科学计数法，例如 -1234.456E+78
%f	有小数点而无指数，例如 123.456
%g	根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
%G	根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
字符串和bytes的slice表示：
%s	字符串或切片的无解译字节
%q	双引号围绕的字符串，由Go语法安全地转义
%x	十六进制，小写字母，每字节两个字符
%X	十六进制，大写字母，每字节两个字符
指针：

%p	十六进制表示，前缀 0x
！

*/
func main() {
	counts := make(map[rune]int)    // unicode字符数量
	var utflen [utf8.UTFMax + 1]int // utf8编码长度
	invalid := 0                    // 非法uff-8字符数量

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbyte, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charCount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid ++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n) // count, number
	}
	fmt.Printf("len\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n) // count, number
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
