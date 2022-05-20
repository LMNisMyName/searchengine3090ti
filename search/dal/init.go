package dal

import (
	"search/dal/db"
)

func Init() {
	db.Init()
}

// func addAndQuery() {
// 	db.Init()
// 	db.AddInvertedIndex(context.Background(), "shanghai", 1)
// 	ids, find := db.Query(context.Background(), "shanghai")
// 	fmt.Println(ids, find)
// }
