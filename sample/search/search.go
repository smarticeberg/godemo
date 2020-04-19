package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

// 执行搜索逻辑
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		// Fatal函数接收错误值，将错误输出，并终止程序
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个wait group，以便处理所有的数据源，防止程序在全部搜索执行完之前终止
	// waitGroup是计数信号量，可以用于统计所有的goroutine是不是都完成了工作
	var waitGroup sync.WaitGroup

	//设置需要等待处理每个数据源的goroutine数量
	waitGroup.Add(len(feeds))

	// 对每个数据源启动一个goroutine来查找结果
	// ‘_’为占位符，占据保存range调用返回的索引值的变量的位置，如果要调用的函数返回多个值，又不需要其中的某个值，可以使用下划线将其忽略掉
	for _, feed := range feeds {
		// 赋值给两个变量时，第一个表示查找的结果值，第二个表示查找的key是否存在于map中
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个goroutine来执行搜索
		// 匿名函数
		// matcher和feed作为变量的值被传入匿名函数
		// 因为matcher、feed每次调用时值不相同，所以并没有使用闭包的方式访问这两个变量，waitGroup,searchTerm，results则使用了闭包的方式
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//启动一个goroutine来监控是否所有的工作都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()
		// 用关闭通道的方式，通知Display函数可以退出程序了
		close(results)
	}()

	// 启动函数，显示返回的结果，并且在最后一个结果显示完后返回
	Display(results)

}

// 注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {

	if _, exists := matchers[feedType]; exists {
		log.Fatal(feedType, "Matcher already registered.")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
