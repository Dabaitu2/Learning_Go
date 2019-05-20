package outline2

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var depth int

func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

// %*s 中的*输出带有可变数量空格的字符串, 宽度和字符串由depth*2 和 ""来提供
func StartElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func EndElement(n *html.Node)  {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	ForEachNode(doc, StartElement, EndElement)
}

