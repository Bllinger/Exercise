package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

var bodyMap = make(map[string]*bytes.Buffer)

func main() {
	fetchAll()

	for url, reader := range bodyMap {
		fmt.Printf("%s中的链接有：\n", url)
		doc, err := html.Parse(reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		for _, link := range visit(nil, doc) {
			fmt.Println("link: ", link)
		}
		reader.Reset()
	}

}

func visit(links []string, n *html.Node) []string {
	//fmt.Println("into visit")
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, e := range n.Attr {
			if e.Key == "href" {
				links = append(links, e.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func fetchAll() {
	ch := make(chan *bytes.Buffer)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for _, url := range os.Args[1:] {
		bodyMap[url] = <-ch
	}
}

func fetch(url string, ch chan *bytes.Buffer) error {
	read := new(bytes.Buffer)

	//if url == "www.sina.com" {
	//	time.Sleep(time.Duration(10)*time.Second)
	//}

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
		fmt.Println("add url success: ", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetch link failed: %v", resp.Status)
	}

	_, err = io.Copy(read, resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	ch <- read
	return nil
	//for _, link := range os.Args[1:] {
	//	if !strings.HasPrefix(link, "http://") {
	//		link = "http://" + link
	//	}
	//
	//	resp, err := http.Get(link)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "%v\n", err)
	//		resp.Body.Close()
	//		return
	//	}
	//
	//	if resp.StatusCode != http.StatusOK {
	//		fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("fetch links failed: %v", resp.Status))
	//		resp.Body.Close()
	//		return
	//	}
	//
	//	_,err = io.Copy(reader, resp.Body)
	//	resp.Body.Close()
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "%v\n", err)
	//		return
	//	}
	//}
}
