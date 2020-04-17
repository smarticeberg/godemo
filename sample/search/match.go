package search

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
