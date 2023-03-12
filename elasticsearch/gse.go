package main

import (
	"fmt"
	"regexp"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	text = "Hello world, Helloworld. Winter is coming! 你好世界."

	new, _ = gse.New("zh,testdata/test_dict3.txt", "alpha")

	seg    gse.Segmenter
	posSeg pos.Segmenter
)

func main() {
	// Loading the default dictionary
	seg.LoadDict()
	// Loading the default dictionary with embed
	// seg.LoadDictEmbed()
	//
	// Loading the Simplified Chinese dictionary
	// seg.LoadDict("zh_s")
	// seg.LoadDictEmbed("zh_s")
	//
	// Loading the Traditional Chinese dictionary
	// seg.LoadDict("zh_t")
	//
	// Loading the Japanese dictionary
	// seg.LoadDict("jp")
	//
	// Load the dictionary
	// seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	cut()

	segCut()
}

func cut() {
	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)
	fmt.Println("analyze: ", new.Analyze(hmm, text))

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)

	reg := regexp.MustCompile(`(\d+年|\d+月|\d+日|[\p{Latin}]+|[\p{Hangul}]+|\d+\.\d+|[a-zA-Z0-9]+)`)
	text1 := `헬로월드 헬로 서울, 2021年09月10日, 3.14`
	hmm = seg.CutDAG(text1, reg)
	fmt.Println("Cut with hmm and regexp: ", hmm, hmm[0], hmm[6])
}

func analyzeAndTrim(cut []string) {
	a := seg.Analyze(cut, "")
	fmt.Println("analyze the segment: ", a)

	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))
}

func cutPos() {
	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)

	posSeg.WithGse(seg)
	po = posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimWithPos(po, "zg")
	fmt.Println("trim pos: ", po)
}

func segCut() {
	// Text Segmentation
	tb := []byte(text)
	fmt.Println(seg.String(text, true))

	segments := seg.Segment(tb)
	// Handle word segmentation results, search mode
	fmt.Println(gse.ToString(segments, true))
}
