package main

import (
	"awesomeProject/ch5/links"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	//for _, url := range os.Args[1:] {
	//	links, err := findLinks(url)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
	//		continue
	//	}
	//	for _, link := range links {
	//		fmt.Println(link)
	//	}
	//}
	breadthFirst(crawl, os.Args[1:])
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/**
	对worklist中的每个元素调用f,j将返回的内容添加到worklist中，每个元素最多调用一次f
 */
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// f(item)会把返回的列表中的所有项展开添加
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
