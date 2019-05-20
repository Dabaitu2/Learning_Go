package main

import (
	"awesomeProject/ch2/popCount"
	"awesomeProject/ch2/tempconv"
	"fmt"
)

func main()  {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZero)
	var a uint64 = 65535
	fmt.Print("[")
	fmt.Print(byte(a >> (7*8)))
	fmt.Print(",")
	fmt.Print(byte(a >> (6*8)))
	fmt.Print(",")
	fmt.Print(byte(a >> (5*8)))
	fmt.Print(",")
	fmt.Print(byte(a >> (4*8)))
	fmt.Print(",")
	fmt.Print(byte(a >> (3*8)))
	fmt.Print(",")
	fmt.Print(byte(a >> (2*8)))
	fmt.Print(",")
	fmt.Print(byte(a >> (1*8)))
	fmt.Print(",")
	fmt.Print(byte(a >> (0*8)))
	fmt.Println("]")
	fmt.Println(popCount.FasterPopCount(a))
	fmt.Print(byte(15>>1))
}
