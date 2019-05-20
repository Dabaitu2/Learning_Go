package main

import (
	"awesomeProject/retriver/Real"
	"awesomeProject/retriver/mock"
	"fmt"
)

// Go语言中的接口，更多是一种辅助描述，表明其是什么
// 而java中的接口类似一种规范，接口是一套规范，而某个类要有功能需要实现这个规范，规范针对接口而开发
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func Post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "tomoko",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}

// 审查一个Retriever, 根据不同的Retriever实际类型打印出不同的内容
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	// 可以用这个语法查询变量的类型(类型是指向某个type的指针)
	// *Type 指的是指针类型定义
	// *variable 指的是取的指针变量对应的实际值
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *Real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	// 接口类型
	var r Retriever
	// 新建一个结构体
	// 这个地方的mockRetriever还不是一个常规实例化的结构体
	// 只是一个字面量值
	// 如果要常规实例化一个结构体，应该使用new
	mockRetriever := mock.Retriever{
		Contents: "this is a fake baidu.com",
	}
	// 获得引用类型的地址
	// 这个结构体实现了这个接口，所以可以把它的地址赋给接口
	// 由于前面没有用new 所以这里需要传地址，即把引用传给这个变量
	r = &mockRetriever
	// 传入的是地址，也就是说兼容于指针变量，
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println("Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
}
