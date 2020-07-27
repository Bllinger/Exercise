package main

import (
	"fmt"
	"main/github"
	"os"
	"time"
)

func main() {
	var during = -1

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	for during != 0 {
		fmt.Println("请输入数字，选择查看多久的issue:")
		fmt.Println("1:一个月以内")
		fmt.Println("2:一年以内")
		fmt.Println("3:超过一年")
		fmt.Println("0:退出")
		fmt.Scanln(&during)

		var issues []*github.Issue

		total, i := 0, 0

		switch during {
		case 1:
			for _, item := range result.Items {
				if daysAgo(item.CreatedAt) <= 30 {
					issues = append(issues, item)
					total++
					i++
				}
			}
		case 2:
			for _, item := range result.Items {
				if daysAgo(item.CreatedAt) <= 365 {
					issues = append(issues, item)
					total++
					i++
				}
			}
		case 3:
			for _, item := range result.Items {
				if daysAgo(item.CreatedAt) > 365 {
					issues = append(issues, item)
					total++
					i++
				}
			}
		case 0:
			fmt.Println("程序结束")
			return

		default:
			fmt.Println("输入错误！！！")
			continue
		}

		items := new(github.IssuesSearchResult)
		*items = github.IssuesSearchResult{TotalCount: total, Items: issues}

		fmt.Printf("Total issue is: %d\n", items.TotalCount)

		for _, item := range items.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
