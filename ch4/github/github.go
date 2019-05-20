package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

/**
	对于一个函数（或方法），如果函数的参数（或接收者）是对象指针时，
	下表示此对象是可被修改的；相反的，如果是对象时，表示是不可修改的
    （但如果该对象本身就是引用类型，如 map\func\chan 等，则本质上是可以修改的）。
	所以一般的做法是，方法的接收者习惯性使用对象指针，而不是对象，一方面可以在想修改对象时进行修改，另一方面也减少参数传递的拷贝成本。
	1.如果你的struct足够大，使用指针可以加快效率
	2.如果不使用指针，在函数内部则无法修改struct中的值
 */
type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue // 这样才能修改其子结构体
}

// json模板字符串制定了marshal之后某字段对应的json
type Issue struct {
	Number   int
	HtmlURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"create_at"`
	Body     string
}

type User struct {
	Login   string
	HtmlURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	// 将请求转译后转换为安全的查询字符串
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		// %-5d 左对齐展示五位数字
		// %9.9s 最小宽度为9, 最大宽度为9并右对齐
		// %.55s 最小宽度为0, 最大宽度为55且右对齐
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
