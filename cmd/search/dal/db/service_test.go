package db_test

import (
	"context"
	"fmt"
	"os"
	"searchengine3090ti/cmd/search/dal/db"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"testing"
)

func TestMain(m *testing.M) {
	db.Init()
	db.DropAllTable()
	// db.CreateTable()
	code := m.Run()
	os.Exit(code)
}

func TestAddIndex(t *testing.T) {
	req := searchapi.AddRequest{Id: 2, Text: "shanghai", Url: "www.shanghai"}
	keywords := []string{"shanghai"}
	db.AddIndex(context.Background(), &req, keywords)
}

//通过关键词查询id
func TestQuery(t *testing.T) {
	ids, find := db.Query(context.Background(), "shanghai")
	fmt.Println(ids)
	fmt.Println(find)
	// assert.Equal(t, true, find)
}

//通过索引id数组查询索引内容(Id, Text, Url)数组
func TestQueryRecord(t *testing.T) {
	ids, _ := db.Query(context.Background(), "extra")
	records, _ := db.QueryRecord(context.Background(), ids)
	fmt.Println(ids)
	fmt.Println(records)
}

//通过id查询关键词
func TestQueryKeyWords(t *testing.T) {
	keywords, find := db.QueryKeyWords(context.Background(), 1)
	fmt.Println(keywords)
	fmt.Println(find)
	// assert.Equal(t, true, find)
}

//通过关键词查询图片id
func TestQueryImagesRecord(t *testing.T) {
	ids, _ := db.Query(context.Background(), "extra")
	records, _ := db.QueryImagesRecord(context.Background(), ids)
	fmt.Println(ids)
	fmt.Println(records)
}
