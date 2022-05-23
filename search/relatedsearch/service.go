package relatedsearch

import (
	"github.com/liyue201/gostl/ds/queue"
	"github.com/mozillazg/go-pinyin"
)

//提供英文(小写)相关搜索服务的字典树节点
type TrieNode struct {
	sons   [26]*TrieNode
	isLeaf bool
	text   string //小写英文
	weight int
}

//提供中文相关搜索服务的字典树节点
type TrieNodeCh struct {
	sons   [26]*TrieNodeCh
	isLeaf bool
	text   string //中文
	weight int
}

type SearchTrie struct {
	root   *TrieNode
	rootCh *TrieNodeCh
}

var searchTrie *SearchTrie

//构造一颗提供搜索字典树
func NewTrie() *SearchTrie {
	newRoot := new(TrieNode)
	newRootCh := new(TrieNodeCh)
	newTrie := SearchTrie{root: newRoot, rootCh: newRootCh}
	return &newTrie
}

func Init() {
	if searchTrie == nil {
		searchTrie = NewTrie()
	}
}

//添加英文内容
func AddEn(text string) {
	curNode := searchTrie.root
	for _, c := range text {
		i := c - 'a'
		if curNode.sons[i] == nil {
			curNode.sons[i] = new(TrieNode)
		}
		curNode = curNode.sons[i]
	}
	curNode.isLeaf = true
	curNode.text = text
	curNode.weight++
}

func GetPinYin(text string) [][]string {
	a := pinyin.NewArgs()
	ans := pinyin.Pinyin(text, a)
	return ans
}

//添加中文内容
// func AddCh(text string) {

// }

//向搜索字典树当中添加内容
func Add(text string) {
	//1. 判断内容是中文内容还是英文内容

	//2. 分别调用相关的接口
	AddEn(text)
}

//返回所有相关搜索结果且不按照权重进行排序
func SearchAllEn(prefix string) []string {
	ans := []string{}
	curNode := searchTrie.root
	//1. 找到前缀字符串最后一个字母所在的节点
	for _, c := range prefix {
		i := c - 'a'
		if curNode.sons[i] == nil {
			return []string{}
		}
		curNode = curNode.sons[i]
	}

	//2. 以当前节点为根节点遍历所有叶子节点（采用BFS）
	que := queue.New()
	que.Push(curNode)
	for {
		f := que.Pop()
		if ret, _ := f.(*TrieNode); ret.isLeaf {
			ans = append(ans, ret.text)
		}
		for i := 0; i < 26; i++ {
			if ret, _ := f.(*TrieNode); ret.sons[i] != nil {
				que.Push(ret.sons[i])
			}
		}
		if que.Empty() {
			break
		}
	}
	return ans
}

//以输入参数为前缀在搜索字典树中查询所有记录
func SearchAll(prefix string) []string {
	//1. 判断内容是中文内容还是英文内容

	//2. 分别调用相关的接口
	return SearchAllEn(prefix)
}

//以输入参数为前缀在搜索字典树中查询前K条记录
// func SearchTopK(prefix string, k int) []string {

// }
