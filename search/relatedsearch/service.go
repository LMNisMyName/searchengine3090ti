package relatedsearch

import (
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
	root   *TrieNode //提供英文（小写字母）相关搜索
	rootCh *TrieNode //提供中文相关搜索
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

func GetPinYin(text string) string {
	a := pinyin.NewArgs()
	pys := pinyin.Pinyin(text, a)
	var ans string
	for _, py := range pys {
		ans += py[0]
	}
	return ans
}

//根据当前字符串的首字母判断是中文还是英文字符串
func isEn(text string) bool {
	if len(text) == 0 {
		return true
	}
	if v := text[0] - 'a'; v >= 26 {
		return false
	}
	return true
}

//向搜索字典树当中添加内容
//判断内容是中文内容还是英文内容, 分别将其添加到对应的树下
func Add(text string) {
	var curNode *TrieNode
	var textEn string
	if isEn(text) {
		curNode = searchTrie.root
		textEn = text
	} else {
		curNode = searchTrie.rootCh
		textEn = GetPinYin(text)
	}
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

//以输入参数为前缀在搜索字典树中查询所有记录
//判断内容是中文内容还是英文内容,分别调用相关搜索树
//返回所有相关搜索结果且不按照权重进行排序
func SearchAll(prefix string) []string {
	ans := []string{}
	var curNode *TrieNode
	var prefixEn string
	if isEn(prefix) {
		curNode = searchTrie.root
		prefixEn = prefix
	} else {
		curNode = searchTrie.rootCh
		prefixEn = GetPinYin(prefix)
	}
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
	var curNode *TrieNode
	var prefixEn string
	if isEn(prefix) {
		curNode = searchTrie.root
		prefixEn = prefix
	} else {
		curNode = searchTrie.rootCh
		prefixEn = GetPinYin(prefix)
	}
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
	ans := make([]string, pq.Size())
	index := pq.Size() - 1
	for {
		ans[index] = (pq.Pop()).(*TrieNode).text
		index--
		if index == -1 {
			break
		}
	}
	return ans
}
