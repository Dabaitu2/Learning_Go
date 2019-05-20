package main

import (
	"fmt"
)

func testMap() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)
	// 空白初始化
	ages = make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
}

func main() {
	testMap()
}
