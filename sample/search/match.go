package search

import (
	"fmt"
	"log"
)

// Result保存搜索的结果，声明结构类型
type Result struct {
	Field   string
	Content string
}

// 声明接口类型
type Matcher interface {
	// 参数为Feed类型的指针
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResult, err := matcher.Search(feed, searchTerm)

	if err != nil {
		log.Println(err)
		return
	}

	//将结果写入通道
	for _, result := range searchResult {
		results <- result
	}
}

// 从每个单独的goroutine接收到结果后从终端窗口输出
func Display(results chan *Result) {
	for result := range results {
		// 通道会一直被阻塞，直到有结果写入
		// 一旦通道被关闭，for循环就会被终止
		fmt.Println("%s:\n%s\n\n", result.Field, result.Content)
	}
}
