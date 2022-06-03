package tokenizer

import (
	"embed"

	"searchengine3090ti/cmd/search/tokenizer/utils"

	"strings"

	"github.com/wangbin/jiebago"
)

var (
	// embed the file as a file systme, which is useful while embedding multiple files
	//That means you can add other files to the dictonaryFS, but be careful to open it in func NewTokenizer
	//go:embed data/*.txt
	dictonaryFS embed.FS
)

type Tokenizer struct {
	seg jiebago.Segmenter // the segment algorithm from open sourced project jieba(golang version)
	// return value of Tokenizer
	/*	type Dictionary struct {
			total, logTotal float64
			freqMap         map[string]float64
			sync.RWMutex
	}······*/
}

var MyTokenizer *Tokenizer

func NewTokenizer(dictionaryPath string) *Tokenizer {
	file, err := dictonaryFS.Open("data/dictionary.txt")
	if err != nil {
		panic(err)
	}
	utils.ReleaseAssets(file, dictionaryPath)
	tokenizer := &Tokenizer{}
	err = tokenizer.seg.LoadDictionary(dictionaryPath)
	if err != nil {
		panic(err)
	}
	return tokenizer
}
func (t *Tokenizer) Cut(text string) []string {
	//不区分大小写
	text = strings.ToLower(text)
	//移除所有的标点符号
	text = utils.RemovePunctuation(text)
	//移除所有的空格
	text = utils.RemoveSpace(text)

	var wordMap = make(map[string]int)

	resultChan := t.seg.CutForSearch(text, true)
	for {
		w, ok := <-resultChan
		if !ok {
			break
		}
		_, found := wordMap[w]
		if !found {
			//去除重复的词
			wordMap[w] = 1
		}
	}

	var wordsSlice []string
	for k := range wordMap {
		wordsSlice = append(wordsSlice, k)
	}

	return wordsSlice
}
