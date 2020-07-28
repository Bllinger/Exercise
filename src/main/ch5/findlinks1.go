package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fetch()
}

func fetch() {
	for _, link := range os.Args[1:] {
		if !strings.HasPrefix(link, "http://") {
			link = "http://" + link
		}

		resp, err := http.Get(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			resp.Body.Close()
			return
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("fetch links failed: %v", resp.Status))
			resp.Body.Close()
			return
		}

		nBytes, err := io.Copy(os.Stdin, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}

		fmt.Println("网页总字节数：", nBytes)
	}
}
