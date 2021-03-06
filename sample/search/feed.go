package search

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
)

const dataFile = "../data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {

	// open file
	// 获取相对路径的文件信息
	_, filename, _, _ := runtime.Caller(1)
	datapath := path.Join(path.Dir(filename), dataFile)
	file, err := os.Open(datapath)
	if err != nil {
		return nil, err
	}

	// read file success,then execute the following operations.
	// defer确保，该函数一定会被调用，哪怕函数意外崩溃终止，也能保证关键字defer安排调用的函数会被执行
	defer file.Close()

	// 当前值为nil的切片，这个切片包含一组指向Feed类型值的指针
	var feeds []*Feed
	// 将文件解码到切片里
	// 这个切片的每一项是一个指向一个Feed类型值的指针，传入切片的地址
	err = json.NewDecoder(file).Decode(&feeds)

	// 不需要检查错误，调用者会做这件事
	return feeds, err
}
