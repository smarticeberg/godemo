package search

// 实现了默认的匹配器
// 空结构在创建实例时，不会分配任何内存
type defaultMatcher struct {
}

// 将默认匹配器注册到程序中
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// search实现了默认的匹配器行为
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
