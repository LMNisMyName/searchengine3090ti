namespace go SearchApi

struct AddRequest {
    1: i32     id
    2: string  text
    3: string  url
}

struct AddResponse {
    1: bool  status 
}

struct QueryRequest {
  1: string queryText
  2: i32    page
  3: i32    limit
  4: i32    order
}

struct QueryResponse {
  1: double time
  2: i32    total
  3: i32    pagecount
  4: i32    page
  5: i32    limit
  6: list<AddRequest>      contents
}

service Search {
    QueryResponse query(1: QueryRequest req)
    AddResponse add(1: AddRequest req)
}