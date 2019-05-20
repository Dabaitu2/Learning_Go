package main

import (
	"awesomeProject/ch5/outline2"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// 检查content-type是不是html
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html; charset=utf-8") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url ,err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	outline2.ForEachNode(doc, visitNode, nil)
	return nil
}

func main()  {
	for _, url := range os.Args[1:] {
		title(url)
	}
}
