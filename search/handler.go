package main

import (
	"context"
	"fmt"
	"math"
	"search/dal/db"
	searchapi "search/kitex_gen/SearchApi"
	"search/relatedsearch"
	"search/tokenizer"
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
	idMap := make(map[int32]int32) //记录每个ID及其出现的次数
	//1.调用两次数据库的查询id服务，分别查询要搜索的id和要屏蔽的id，建立id与其出现次数的映射
	queryKeywords := tokenizer.MyTokenizer.Cut(req.QueryText)
	filterKeyWords := tokenizer.MyTokenizer.Cut(req.FilterText)
	for _, qword := range queryKeywords {
		if ids, find := db.Query(ctx, qword); find {
			for _, id := range ids {
				idMap[id] += 1
			}
		}
	}
	for _, fword := range filterKeyWords {
		if ids, find := db.Query(ctx, fword); find {
			for _, id := range ids {
				delete(idMap, id)
			}
		}
	}
	//2. 对id差集按照出现次数（表示关联度）从大到小进行排序
	v := vector.New() //v是保存根据出现次数排序的数组
	for id := range idMap {
		v.PushBack(id)
	}
	//按照idMap中记录的出现次数对id进行排序
	// req.Order == 0 按照次数降序排序(默认)， == 1 按照次数升序排序
	if req.Order == 0 {
		sort.Sort(v.Begin(), v.End(), func(l, r any) int {
			if l == r {
				return 0
			}
			switch l.(type) {
			case int32:
				if idMap[l.(int32)] > idMap[r.(int32)] {
					return 1
				} else {
					return -1
				}
			}
			return 1
		})
	} else if req.Order == 1 {
		sort.Sort(v.Begin(), v.End(), func(l, r any) int {
			if l == r {
				return 0
			}
			switch l.(type) {
			case int32:
				if idMap[l.(int32)] > idMap[r.(int32)] {
					return -1
				} else {
					return 1
				}
			}
			return 1
		})
	}
	//3. 根据分页请求在排好序的数组中确定一个窗口 [lIndex, rIndex)
	resp = new(searchapi.QueryResponse)
	resp.Page = req.Page
	resp.Limit = req.Limit
	resp.Pagecount = int32(math.Ceil((float64(v.Size()) / float64(req.Limit))))
	resp.Contents = make([]*searchapi.AddRequest, 0)
	lIndex := (req.Page - 1) * req.Limit
	rIndex := req.Page * req.Limit
	if rIndex > int32(v.Size()) {
		rIndex = int32(v.Size())
	}
	if lIndex >= int32(v.Size()) {
		//窗口超过数组上界
		return
	} else {
		resp.Total = rIndex - lIndex
		ids := make([]int32, resp.Total)
		for i := lIndex; i < rIndex; i++ {
			ids[i-lIndex] = v.At(int(i)).(int32)
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

// FindID implements the SearchImpl interface.
func (s *SearchImpl) FindID(ctx context.Context, req *searchapi.FindIDRequest) (resp *searchapi.FindIDResponse, err error) {
	// 调用数据库根据ID查询记录接口
	ret, err := db.QueryRecord(context.Background(), []int32{req.Id})
	if len(ret) == 0 {
		resp.Found = false
	} else {
		resp.Found = true
	}
	return
}
