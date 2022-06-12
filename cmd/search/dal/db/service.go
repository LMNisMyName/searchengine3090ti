package db

import (
	"context"
	"errors"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
	"strings"

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
	UID      int64
	Keywords []*Keyword `gorm:"many2many:keyword_userids;"`
}

type Record struct {
	gorm.Model
	RecordId int64
	Text     string
	Url      string
}

// 对于Record表，直接新增一条记录即可
// 对于KeyWords表，需要添加目前表中尚未存在的keywords
// 同时建立起这些Keywords与UserID记录之间的关联
// 对于UserID表，添加当前记录
// 同时建立起当前UserID与Keywords之间的关联
func AddIndex(ctx context.Context, req *searchapi.AddRequest, keywords []string) error {
	//添加record记录
	newRecordEntry := Record{RecordId: req.Id, Text: req.Text, Url: req.Url}
	if ret := DB.Model(&Record{}).Create(&newRecordEntry); ret.Error != nil {
		return ret.Error
	}
	//添加UserID记录
	var newUserIDEntry UserID
	result := DB.First(&newUserIDEntry, "uid = ?", req.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newUserIDEntry = UserID{UID: req.Id}
		if ret := DB.Model(&UserID{}).Create(&newUserIDEntry); ret.Error != nil {
			return ret.Error
		}
	}
	//遍历关键词添加尚未存在的keywords记录
	keywordSlice := []*Keyword{}
	for _, word := range keywords {
		var keywordEntry Keyword
		//如果当前关键词没有找到那么新建一个
		result := DB.First(&keywordEntry, "word = ?", word)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			keywordEntry = Keyword{Word: word}
			if ret := DB.Model(&Keyword{}).Create(&keywordEntry); ret.Error != nil {
				return ret.Error
			}
		}
		//建立Keywords与UserID记录之间的关联
		if ret := DB.Model(&keywordEntry).Association("UserIDs").Append(&newUserIDEntry); ret != nil {
			return ret
		}
		keywordSlice = append(keywordSlice, &keywordEntry)
	}
	if ret := DB.Model(&newUserIDEntry).Association("Keywords").Append(keywordSlice); ret != nil {
		return ret
	}
	return nil
}

//查询与关键词相关的索引id数组
func Query(ctx context.Context, keyword string) ([]int64, bool) {
	var keywordEntry Keyword
	result := DB.Model(&Keyword{}).First(&keywordEntry, "word = ?", keyword)
	if result.Error != nil {
		return []int64{}, false
	}
	userIds := []UserID{}
	DB.Model(&keywordEntry).Association("UserIDs").Find(&userIds)
	ids := make([]int64, len(userIds))
	for index, userId := range userIds {
		ids[index] = userId.UID
	}
	return ids, true
}

//通过索引id数组查询索引内容(Id, Text, Url)数组
func QueryRecord(ctx context.Context, ids []int64) ([]searchapi.AddRequest, error) {
	ans := make([]searchapi.AddRequest, len(ids))
	var err error
	for i, id := range ids {
		var recordEntry Record
		result := DB.Model(&Record{}).First(&recordEntry, "record_id = ?", id)
		if result.Error != nil {
			err = result.Error
		}
		ans[i] = searchapi.AddRequest{Id: recordEntry.RecordId, Text: recordEntry.Text, Url: recordEntry.Url}
	}
	return ans, err
}

//查询与id相关的关键词
func QueryKeyWords(ctx context.Context, id int64) ([]string, bool) {
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

//查询当前记录的数目
func QueryRecordsNumber(ctx context.Context) (ans int64, err error) {
	var records []Record
	result := DB.Model(&Record{}).Find(&records)
	if result.Error != nil {
		err = result.Error
	}
	if records == nil {
		ans = 0
	} else {
		ans = int64(len(records))
	}
	return
}

//通过索引id数组查询索引内容为图片(Id, Text, Url)的数组
func QueryImagesRecord(ctx context.Context, ids []int64) ([]searchapi.AddRequest, error) {
	ans := []searchapi.AddRequest{}
	var err error
	for _, id := range ids {
		var recordEntry Record
		result := DB.Model(&Record{}).First(&recordEntry, "record_id = ?", id)
		if result.Error != nil {
			err = result.Error
		}
		if strings.HasPrefix(recordEntry.Url, "http") && (strings.Contains(recordEntry.Url, ".jpg") ||
			strings.Contains(recordEntry.Url, ".gif") || strings.Contains(recordEntry.Url, ".jpeg") ||
			strings.Contains(recordEntry.Url, ".png")) {
			ans = append(ans, searchapi.AddRequest{Id: recordEntry.RecordId, Text: recordEntry.Text, Url: recordEntry.Url})
		}

	}
	return ans, err
}
