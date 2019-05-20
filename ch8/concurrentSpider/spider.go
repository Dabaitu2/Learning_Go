package main

import (
	"awesomeProject/ch5/links"
	"fmt"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)


func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main()  {
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	worklist := make(chan []string)
	go func() { worklist <- os.Args[1:] }()
	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}