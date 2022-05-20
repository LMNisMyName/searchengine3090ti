package db

import (
	"context"
	"errors"
	searchapi "search/kitex_gen/SearchApi"

	"gorm.io/gorm"
)

type Keyword struct {
	gorm.Model
	Word    string
	UserIDs []*UserID `gorm:"many2many:keyword_userids;"`
}

type UserID struct {
	gorm.Model
	UID      int32
	Keywords []*Keyword `gorm:"many2many:keyword_userids;"`
}

type Record struct {
	gorm.Model
	RecordId int32
	Text     string
	Url      string
}

// 要求函数的调用方提供的keyword和id不能同时重复
func AddIndex(ctx context.Context, req *searchapi.AddRequest, keywords []string) {
	newRecordEntry := Record{RecordId: req.Id, Text: req.Text, Url: req.Url}
	DB.Model(&Record{}).Create(&newRecordEntry)
	newUserIDEntry := UserID{UID: req.Id}
	DB.Model(&UserID{}).Create(&newUserIDEntry)
	keywordSlice := []*Keyword{}
	for _, word := range keywords {
		var keywordEntry Keyword
		//如果当前关键词没有找到那么新建一个
		result := DB.First(&keywordEntry, "word = ?", word)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			keywordEntry := Keyword{Word: word}
			DB.Model(&Keyword{}).Create(&keywordEntry)
		}
		DB.Model(&keywordEntry).Association("UserIDs").Append(&newUserIDEntry)
		keywordSlice = append(keywordSlice, &keywordEntry)
	}
	DB.Model(&newUserIDEntry).Association("Keywords").Append(keywordSlice)
}

//利用倒排索引数据库查询与关键词相关的索引id数组
func Query(ctx context.Context, keyword string) ([]int32, bool) {
	var keywordEntry Keyword
	result := DB.First(&keywordEntry, "word = ?", keyword)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return []int32{}, false
	// }
	if result.Error != nil {
		return []int32{}, false
	}
	idPointers := []*UserID{}
	DB.Model(&keywordEntry).Association("UserIDs").Find(&idPointers)
	ids := make([]int32, len(idPointers))
	for index, idPointer := range idPointers {
		ids[index] = idPointer.UID
	}
	return ids, true
}
