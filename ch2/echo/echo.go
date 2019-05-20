package main

// flag包可以解析命令行参数
import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
