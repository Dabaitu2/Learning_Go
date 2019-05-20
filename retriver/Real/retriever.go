package Real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {

	// 通过url获取响应对象
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// 倾倒响应结果
	result, err := httputil.DumpResponse(
		resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
