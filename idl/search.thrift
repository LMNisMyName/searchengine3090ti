namespace go SearchApi

struct AddRequest {
    1: i64     id       //id序号（前端要尽量保证其不重复，最好是升序）
    2: string  text     //描述图片的相关文本
    3: string  url      //图片的url链接
}

struct AddResponse {
    1: bool  status     //数据库添加是否成功
}

struct QueryRequest {
  1: string queryText   //用户在搜索框中输入的要查询的关键词
  2: string filterText  //用户请求过滤的关键词
  3: i64    page        //用户请求的页码
  4: i64    limit       //每页显示的请求数目
  5: i64    order       //排序方式
}

struct QueryResponse {
  1: double time        //查询所需时间
  2: i64    total       //查询到的条目总数 
  3: i64    pagecount   //查询到的页数 
  4: i64    page        //当前页码 
  5: i64    limit       //每页展示的数目  
  6: list<AddRequest>      contents //查询到的内容
}
struct imgQueryRequest {
  1: string queryText   //用户在搜索框中输入的要查询的关键词
  2: string filterText  //用户请求过滤的关键词
  3: i64    page        //用户请求的页码
  4: i64    limit       //每页显示的请求数目
  5: i64    order       //排序方式
}

struct imgQueryResponse {
  1: double time        //查询所需时间
  2: i64    total       //查询到的条目总数
  3: i64    pagecount   //查询到的页数
  4: i64    page        //当前页码
  5: i64    limit       //每页展示的数目
  6: list<AddRequest>      contents //查询到的内容
}
struct wd2imgQueryRequest {
  1: string queryText   //用户在搜索框中输入的要查询的关键词
  2: string filterText  //用户请求过滤的关键词
  3: i64    page        //用户请求的页码
  4: i64    limit       //每页显示的请求数目
  5: i64    order       //排序方式
}

struct wd2imgQueryResponse {
  1: double time        //查询所需时间
  2: i64    total       //查询到的条目总数
  3: i64    pagecount   //查询到的页数
  4: i64    page        //当前页码
  5: i64    limit       //每页展示的数目
  6: list<AddRequest>      contents //查询到的内容
}
struct RelatedQueryRequest{
    1: string queryText             //用户在搜索框中输入的要查询的关键词
}

struct RelatedQueryResponse{
    1: list<string> relatedTexts    //与用户输入请求相关的文本
}

struct FindIDResponse{
    1: list<AddRequest> entries                 //是否找到
}

struct FindIDRequest{
    1: list<i64>  ids                      //要查找的ID   
}

struct QueryIDNumberResponse {
    1 : i64 number
}

struct QueryIDNumberRequest{
}


service Search {
    //提供支持分页、关键词过滤的查询服务
    QueryResponse query(1: QueryRequest req)
    //提供添加索引服务
    AddResponse add(1: AddRequest req)
    //提供相关搜索服务
    RelatedQueryResponse relatedQuery(1: RelatedQueryRequest req)
    //查询id是否存在
    FindIDResponse  findID(1: FindIDRequest req) 
    //查询当前记录数目
    QueryIDNumberResponse queryIDNumber(1 : QueryIDNumberRequest req)
    //提供以图搜图查询服务
    imgQueryResponse imgquery(1: imgQueryRequest req)
    //提供文字搜图查询服务
    wd2imgQueryResponse wd2imgquery(1: wd2imgQueryRequest req)
}