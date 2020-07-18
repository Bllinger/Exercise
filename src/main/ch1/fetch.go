package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			fmt.Println("添加成功！")
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "网络请求失败，错误信息为：%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("状态码为：%v\n", resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "获取请求体失败，错误信息为：%v\n", err)
			os.Exit(1)
		}

		//fmt.Println()

		//content,err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//
		//if err != nil {
		//	fmt.Fprintf(os.Stderr,"获取请求体失败，错误信息为：%v\n",err)
		//	os.Exit(1)
		//}
		//
		//fmt.Println(string(content))
	}
}
