package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {

	// open file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// read file success,then execute the following operations.
	// defer确保，该函数一定会被调用，哪怕函数意外崩溃终止，也能保证关键字defer安排调用的函数会被执行
	defer file.Close()

	var feeds []*Feed
	// 将文件解码到切片里
	// 这个切片的每一项是一个指向一个Feed类型值的指针
	err = json.NewDecoder(file).Decode(&feeds)

	// 不需要检查错误，调用者会做这件事
	return feeds, err
}
