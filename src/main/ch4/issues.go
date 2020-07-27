package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var during = -1

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	for during != 0 {
		fmt.Println("请输入数字选择查看多久的issue:")
		fmt.Println("1:一个月以内")
		fmt.Println("2:一年以内")
		fmt.Println("3:超过一年")
		fmt.Println("0:退出")
		fmt.Scanln(&during)

		issues := []Issue{*result.Items[0]}

		total, i := 0, 0

		switch during {
		case 1:
			for _, item := range result.Items {
				if daysAgo(item.CreatedAt) <= 30 {
					issues = append(issues, *item)
					total++
					i++
				}
			}
		case 2:
			for _, item := range result.Items {
				if daysAgo(item.CreatedAt) <= 365 {
					issues = append(issues, *item)
					total++
					i++
				}
			}
		case 3:
			for _, item := range result.Items {
				if daysAgo(item.CreatedAt) > 365 {
					issues = append(issues, *item)
					total++
					i++
				}
			}
		default:
			fmt.Println("输入错误！！！")
		}

		var items = IssuesSearchResult{TotalCount: total, Items: nil}
		items.Items = &issues

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
