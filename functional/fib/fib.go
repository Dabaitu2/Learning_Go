package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 指定一个intGen类型为返回int的函数，这个函数类型实现了Reader接口，拥有read方法
type intGen func() int

// 函数名为Read,是intGen的方法，参数为byte数组，返回值为int和一个错误对象
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	// 需要指定终止条件
	if next > 10000 {
		return 0, io.EOF
	}
	// 格式化某个输入
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	// 把实际产生read返回值的工作代理给其他reader处理
	// 这个函数实际上没有真的重写reader
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	// ioScanner 需要接上一个 reader
	scanner := bufio.NewScanner(reader)
	// go语言中没有while, 用for来替代， 用法和正常的while相同
	// scan除非遇到EOF才停止
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	fib := fibonacci()
	printFileContents(fib)
}
