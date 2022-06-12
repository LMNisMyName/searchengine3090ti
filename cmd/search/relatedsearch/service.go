package relatedsearch

import (
	"bytes"

	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/ds/queue"
	"github.com/mozillazg/go-pinyin"
)

//提供相关搜索服务的字典树节点
type TrieNode struct {
	sons   [26]*TrieNode
	isLeaf bool
	text   string //小写英文
	weight int
}

type SearchTrie struct {
	root   *TrieNode //提供纯英文相关搜索
	rootCh *TrieNode //提供中英文混合相关搜索
}

var searchTrie *SearchTrie

//构造两棵提供中英文相关搜索的字典树
func NewTrie() *SearchTrie {
	newRoot := new(TrieNode)
	newRootCh := new(TrieNode)
	newTrie := SearchTrie{root: newRoot, rootCh: newRootCh}
	return &newTrie
}

func Init() {
	if searchTrie == nil {
		searchTrie = NewTrie()
	}
}

//将纯中文字符串转化为汉语拼音（小写英文字符）
func GetPinYin(text string) string {
	a := pinyin.NewArgs()
	pys := pinyin.Pinyin(text, a)
	var ans string
	for _, py := range pys {
		ans += py[0]
	}
	return ans
}

//将中英文混合字符串转化为小写英文字符串
func GetPinYinForMix(text string) string {
	var b bytes.Buffer
	for _, c := range text {
		if c > 127 {
			b.WriteString(GetPinYin(string(c)))
		} else {
			if c >= 65 && c <= 90 {
				b.WriteRune(c - 65 + 97)
			} else if c >= 97 && c <= 122 {
				b.WriteRune(c)
			}
		}
	}
	return b.String()
}

//根据当前字符串的首字母判断是纯英文还是中英文字符串
func isEn(text string) bool {
	for _, c := range text {
		if c > 127 {
			return false
		}
	}
	return true
}

//向搜索字典树当中添加内容
//判断内容是中文内容还是英文内容, 分别将其添加到对应的树下
func Add(text string) {
	var curNode *TrieNode
	if isEn(text) {
		curNode = searchTrie.root
	} else {
		curNode = searchTrie.rootCh
	}
	textEn := GetPinYinForMix(text)
	for _, c := range textEn {
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

func SearchInCurNode(prefixEn string, curNode *TrieNode) []string {
	ans := []string{}
	//1. 找到前缀字符串最后一个字母所在的节点
	for _, c := range prefixEn {
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

func SearchTopKInCurNode(pq *priorityqueue.PriorityQueue, k int, prefixEn string, curNode *TrieNode) {
	//1. 找到前缀字符串最后一个字母所在的节点
	for _, c := range prefixEn {
		i := c - 'a'
		if curNode.sons[i] == nil {
			return
		}
		curNode = curNode.sons[i]
	}
	//2. 以当前节点为根节点遍历所有叶子节点（采用BFS）
	que := queue.New()
	que.Push(curNode)
	for {
		f := que.Pop()
		if ret, _ := f.(*TrieNode); ret.isLeaf {
			if pq.Size() == k {
				if ret.weight > (pq.Top().(*TrieNode)).weight {
					pq.Pop()
					pq.Push(ret)
				}
			} else {
				pq.Push(ret)
			}
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
}

//以输入参数为前缀在两颗字典树中查询所有记录
//返回所有相关搜索结果且不按照权重进行排序
func SearchAll(prefix string) []string {
	ans := []string{}
	prefixEn := GetPinYinForMix(prefix)
	ans = append(ans, SearchInCurNode(prefixEn, searchTrie.root)...)
	ans = append(ans, SearchInCurNode(prefixEn, searchTrie.rootCh)...)
	return ans
}

//以输入参数为前缀在搜索字典树中查询前K条记录
//建立一个小根堆，如果堆内元素小于k个直接入堆
//否则与堆顶元素进行对比，大入堆
func SearchTopK(prefix string, k int) []string {
	pq := priorityqueue.New(priorityqueue.WithComparator(func(l, r any) int {
		if l == r {
			return 0
		}
		switch l.(type) {
		case *TrieNode:
			if (l.(*TrieNode)).weight > (r.(*TrieNode)).weight {
				return 1
			}
			return -1
		}
		return 1
	}))
	prefixEn := GetPinYinForMix(prefix)
	//查询的字符串为空
	if len(prefixEn) == 0 {
		return []string{}
	}
	SearchTopKInCurNode(pq, k, prefixEn, searchTrie.root)
	SearchTopKInCurNode(pq, k, prefixEn, searchTrie.rootCh)
	ans := make([]string, pq.Size())
	index := pq.Size() - 1
	if index < 0 {
		return ans
	}
	for {
		ans[index] = (pq.Pop()).(*TrieNode).text
		index--
		if index == -1 {
			break
		}
	}
	return ans
}
