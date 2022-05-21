package db

import (
	"context"
	"errors"
	searchapi "search/kitex_gen/SearchApi"

	"gorm.io/gorm"
)

//数据模型
//Keyword与UserID之间是manytomany关系
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

// 对于Record表，直接新增一条记录即可
// 对于KeyWords表，需要添加目前表中尚未存在的keywords
// 同时建立起这些Keywords与UserID记录之间的关联
// 对于UserID表，添加当前记录
// 同时建立起当前UserID与Keywords之间的关联
func AddIndex(ctx context.Context, req *searchapi.AddRequest, keywords []string) {
	//添加record记录
	newRecordEntry := Record{RecordId: req.Id, Text: req.Text, Url: req.Url}
	DB.Model(&Record{}).Create(&newRecordEntry)
	//添加UserID记录
	var newUserIDEntry UserID
	result := DB.First(&newUserIDEntry, "uid = ?", req.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newUserIDEntry = UserID{UID: req.Id}
		DB.Model(&UserID{}).Create(&newUserIDEntry)
	}
	//遍历关键词添加尚未存在的keywords记录
	keywordSlice := []*Keyword{}
	for _, word := range keywords {
		var keywordEntry Keyword
		//如果当前关键词没有找到那么新建一个
		result := DB.First(&keywordEntry, "word = ?", word)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			keywordEntry = Keyword{Word: word}
			DB.Model(&Keyword{}).Create(&keywordEntry)
		}
		//建立Keywords与UserID记录之间的关联
		DB.Model(&keywordEntry).Association("UserIDs").Append(&newUserIDEntry)
		keywordSlice = append(keywordSlice, &keywordEntry)
	}
	DB.Model(&newUserIDEntry).Association("Keywords").Append(keywordSlice)
}

//查询与关键词相关的索引id数组
func Query(ctx context.Context, keyword string) ([]int32, bool) {
	var keywordEntry Keyword
	result := DB.Model(&Keyword{}).First(&keywordEntry, "word = ?", keyword)
	if result.Error != nil {
		return []int32{}, false
	}
	userIds := []UserID{}
	DB.Model(&keywordEntry).Association("UserIDs").Find(&userIds)
	ids := make([]int32, len(userIds))
	for index, userId := range userIds {
		ids[index] = userId.UID
	}
	return ids, true
}

//查询与id相关的关键词
func QueryKeyWords(ctx context.Context, id int32) ([]string, bool) {
	var UserIDEntry UserID
	result := DB.Model(&UserID{}).First(&UserIDEntry, "uid = ?", id)
	if result.Error != nil {
		return []string{}, false
	}
	keywords := []Keyword{}
	DB.Model(&UserIDEntry).Association("Keywords").Find(&keywords)
	words := make([]string, len(keywords))
	for index, keyword := range keywords {
		words[index] = keyword.Word
	}
	return words, true
}
