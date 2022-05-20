package main

import (
	"context"
	searchapi "search/kitex_gen/SearchApi"
)

// SearchImpl implements the last service interface defined in the IDL.
type SearchImpl struct{}

// Query implements the SearchImpl interface.
func (s *SearchImpl) Query(ctx context.Context, req *searchapi.QueryRequest) (resp *searchapi.QueryResponse, err error) {
	// TODO: Your code here...
	return
}

// Add implements the SearchImpl interface.
func (s *SearchImpl) Add(ctx context.Context, req *searchapi.AddRequest) (resp *searchapi.AddResponse, err error) {
	// TODO: Your code here...
	return
}
