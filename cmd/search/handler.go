package main

import (
	"context"
	"fmt"
	"math"
	"searchengine3090ti/cmd/search/dal/db"
	"searchengine3090ti/cmd/search/relatedsearch"
	"searchengine3090ti/cmd/search/tokenizer"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"sync"
	"time"

	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/vector"
)

// SearchImpl implements the last service interface defined in the IDL.
type SearchImpl struct{}

// Query implements the SearchImpl interface.
// 0.将当前查询添加到字典树当中
// 1.调用两次数据库的查询id服务，分别查询要搜索的id和要屏蔽的id，建立id与其出现次数的映射
// 2.按照出现次数（表示关联度）从大到小对id进行排序
// 3.根据分页请求在排好序的数组中确定一个窗口
// 4.对上一步确定的id窗口调用数据库根据id查询记录服务
func (s *SearchImpl) Query(ctx context.Context, req *searchapi.QueryRequest) (resp *searchapi.QueryResponse, err error) {
	startTime := time.Now()
	defer func() {
		endTime := time.Now()
		resp.Time = float64(endTime.Local().UnixMilli()) - float64(startTime.Local().UnixMilli())
	}()
	//0.将当前查询添加到字典树当中
	relatedsearch.Add(req.QueryText)
	idMap := make(map[int64]float64) //记录每个ID及其出现的次数
	//1.调用两次数据库的查询id服务，分别查询要搜索的id和要屏蔽的id，建立id与其出现次数的映射
	queryKeywords := tokenizer.MyTokenizer.Cut(req.QueryText)
	filterKeyWords := tokenizer.MyTokenizer.Cut(req.FilterText)
	//使用Go协程加速idMap的更新和删除
	//使用锁用来保护idMap，使用waitGroup用来同步协程退出
	var mux sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(len(queryKeywords))
	for _, qword := range queryKeywords {
		go func(qword string) {
			if ids, find := db.Query(ctx, qword); find {
				for _, id := range ids {
					var weight float64
					weight, err = getWeight(qword)
					mux.Lock()
					idMap[id] += weight
					mux.Unlock()
				}
			}
			wg.Done()
		}(qword)
	}
	wg.Wait()
	wg.Add(len(filterKeyWords))
	for _, fword := range filterKeyWords {
		go func(fword string) {
			if ids, find := db.Query(ctx, fword); find {
				for _, id := range ids {
					mux.Lock()
					delete(idMap, id)
					mux.Unlock()
				}
			}
			wg.Done()
		}(fword)
	}
	wg.Wait()
	//2. 对id差集按照关联度从大到小进行排序
	v := vector.New()
	for id := range idMap {
		v.PushBack(id)
	}
	// 排序方式
	// req.Order == 0 表示按照关联度从高到低排序
	if req.Order == 0 {
		sort.Sort(v.Begin(), v.End(), func(l, r any) int {
			if l == r {
				return 0
			}
			switch l.(type) {
			case int64:
				if idMap[l.(int64)] > idMap[r.(int64)] {
					return 1
				} else {
					return -1
				}
			}
			return 1
		})
	}
	//3. 根据分页请求在排好序的数组中确定一个窗口 [lIndex, rIndex)
	resp = new(searchapi.QueryResponse)
	resp.Page = req.Page
	resp.Limit = req.Limit
	resp.Pagecount = int64(math.Ceil((float64(v.Size()) / float64(req.Limit))))
	resp.Contents = make([]*searchapi.AddRequest, 0)
	lIndex := (req.Page - 1) * req.Limit
	rIndex := req.Page * req.Limit
	if rIndex > int64(v.Size()) {
		rIndex = int64(v.Size())
	}
	if lIndex >= int64(v.Size()) {
		//窗口超过数组上界
		return
	} else {
		resp.Total = rIndex - lIndex
		ids := make([]int64, resp.Total)
		for i := lIndex; i < rIndex; i++ {
			ids[i-lIndex] = v.At(int(i)).(int64)
		}
		//4. 调用数据库查询records服务
		records, err := db.QueryRecord(ctx, ids)
		fmt.Println(ids)
		if err == nil {
			for _, record := range records {
				recordNew := record
				resp.Contents = append(resp.Contents, &recordNew)
			}
		}
	}
	return
}

// Add implements the SearchImpl interface.
// 添加索引直接调用分词器和db接口即可(使用前确保数据库和分词器已经初始化)
func (s *SearchImpl) Add(ctx context.Context, req *searchapi.AddRequest) (resp *searchapi.AddResponse, err error) {
	keywords := tokenizer.MyTokenizer.Cut(req.Text)
	err = db.AddIndex(ctx, req, keywords)
	resp = new(searchapi.AddResponse)
	if err == nil {
		resp.Status = true
	} else {
		resp.Status = false
	}
	return
}

// RelatedQuery implements the SearchImpl interface.
// 直接调用相关搜索接口即可
func (s *SearchImpl) RelatedQuery(ctx context.Context, req *searchapi.RelatedQueryRequest) (resp *searchapi.RelatedQueryResponse, err error) {
	// 默认返回前10条热门的相关搜索结果
	resp = new(searchapi.RelatedQueryResponse)
	resp.RelatedTexts = relatedsearch.SearchTopK(req.QueryText, 10)
	return
}

//FindID implements the SearchImpl interface.
func (s *SearchImpl) FindID(ctx context.Context, req *searchapi.FindIDRequest) (resp *searchapi.FindIDResponse, err error) {
	// 调用数据库根据ID查询记录接口
	ret, err := db.QueryRecord(context.Background(), []int64{req.Id})
	resp = new(searchapi.FindIDResponse)
	if len(ret) == 0 {
		resp.Found = false
	} else {
		resp.Found = true
	}
	return
}

// QueryIDNumber implements the SearchImpl interface.
func (s *SearchImpl) QueryIDNumber(ctx context.Context, req *searchapi.QueryIDNumberRequest) (resp *searchapi.QueryIDNumberResponse, err error) {
	resp = new(searchapi.QueryIDNumberResponse)
	resp.Number, err = db.QueryRecordsNumber(ctx)
	return
}

//计算关键词的权重（IDF逆文档频率算法）
//关键词在记录中出现的次数越多，关键词的权重越小
func getWeight(keyword string) (float64, error) {
	ids, find := db.Query(context.Background(), keyword)
	number, err := db.QueryRecordsNumber(context.Background())
	if err != nil {
		return 0, err
	}
	if !find {
		return math.Log(float64(number)), nil
	} else {
		return math.Log(float64(len(ids))/float64(number) + 1), nil
	}
}
