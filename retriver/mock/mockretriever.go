package mock

import "fmt"

// Go 语言中，只要参数是结构体的函数就会被看作这个结构体的成员函数
type Retriever struct {
	Contents string
}

// 在函数名前面写括号并表明其类型是为类型添加方法的惯用手段
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriver: {Contents=%s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["Contents"]
	return "OK"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}

