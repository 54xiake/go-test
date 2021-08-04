package main

import (
	"fmt"
	"github.com/huichen/sego"
	"os"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(os.Getenv("GOPATH") + "/pkg/mod/github.com/huichen/sego@v0.0.0-20180617034105-3f3c8a8cfacc/data/dictionary.txt")

	// 分词
	text := []byte("中华人民共和国中央人民政府")
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	fmt.Println(sego.SegmentsToString(segments, false))
}
