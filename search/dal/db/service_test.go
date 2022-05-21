package db_test

import (
	"context"
	"fmt"
	"os"
	"search/dal/db"
	searchapi "search/kitex_gen/SearchApi"
	"testing"
)

func TestMain(m *testing.M) {
	// db.DropAllTable()
	db.Init()
	// db.DelectAllEntry()
	code := m.Run()
	os.Exit(code)
}

func TestAddIndex(t *testing.T) {
	req := searchapi.AddRequest{Id: 1, Text: "", Url: "www"}
	keywords := []string{"extra"}
	db.AddIndex(context.Background(), &req, keywords)
}

func TestQuery(t *testing.T) {
	ids, find := db.Query(context.Background(), "people")
	fmt.Println(ids)
	fmt.Println(find)
	// assert.Equal(t, true, find)
}

func TestQueryKeyWords(t *testing.T) {
	keywords, find := db.QueryKeyWords(context.Background(), 1)
	fmt.Println(keywords)
	fmt.Println(find)
	// assert.Equal(t, true, find)
}
