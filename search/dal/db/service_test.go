package db_test

import (
	"context"
	"fmt"
	"os"
	"search/dal/db"
	searchapi "search/kitex_gen/SearchApi"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	db.Init()
	code := m.Run()
	os.Exit(code)
}

func TestAddIndex(t *testing.T) {
	req := searchapi.AddRequest{Id: 5, Text: "shanghai", Url: "www"}
	keywords := []string{"shanghai"}
	db.AddIndex(context.Background(), &req, keywords)
}

func TestQuery(t *testing.T) {
	ids, find := db.Query(context.Background(), "shanghai")
	fmt.Println(ids)
	fmt.Println(find)
	// assert.Equal(t, ids[0], 1)
	assert.Equal(t, find, true)
}
