package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _,url := range os.Args[1:]{
		go fetch(url,ch)
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("Total Time:%.2fs\n",time.Since(start).Seconds())
}

func fetch(url string,ch chan string)  {
	start := time.Now()

	if !strings.HasPrefix(url,"http://") {
		url = "http://"+url
		fmt.Printf("%s的http添加成功\n",url)
	}
	
	resp,err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes,err := io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprint(err)
	}

	ch <- fmt.Sprintf("%.2fs\t%d\t%s",time.Since(start).Seconds(),nbytes,url)
}