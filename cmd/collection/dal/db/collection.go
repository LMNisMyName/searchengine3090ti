package db

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"searchengine3090ti/pkg/constants"

	"gorm.io/gorm"
)

type Entry []int64

func (e *Entry) Scan(value interface{}) error {
	byteValue, _ := value.([]byte)
	return json.Unmarshal(byteValue, e)
}

func (e *Entry) Value() (driver.Value, error) {
	return json.Marshal(e)
}

type Collection struct {
	gorm.Model
	UserID int64 `json:"user_id"`

	Name    string `json:"name"`
	Entries Entry  `json:"entries"`
}

func (c *Collection) TableName() string {
	return constants.CollectionTableName
}

//创建收藏夹
func CreateCollection(ctx context.Context, collections []*Collection) error {
	return DB.WithContext(ctx).Create(collections).Error
}

//获取用户收藏夹
func MGetColletction(ctx context.Context, UserID int64) ([]*Collection, error) {
	res := make([]*Collection, 0)
	if err := DB.WithContext(ctx).Where("user_id = ?", UserID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//获取指定收藏夹
func GetColletction(ctx context.Context, UserID, ColltID int64) ([]*Collection, error) {
	res := make([]*Collection, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND id = ?", UserID, uint(ColltID)).Find(&res).First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

//删除收藏夹
func DeleteCollection(ctx context.Context, UserID, ColltID int64) error {
	res := make([]*Collection, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND id = ?", UserID, uint(ColltID)).Delete(&res).Error; err != nil {
		return err
	}
	return nil
}

//添加收藏夹记录
func AddEntry(ctx context.Context, UserID, ColltID int64, newEntry int64) error {
	res := make([]*Collection, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND id = ?", UserID, uint(ColltID)).Find(&res).First(&res).Error; err != nil {
		return err
	}
	res[0].Entries = append(res[0].Entries, newEntry)
	if err := DB.WithContext(ctx).Save(res).Error; err != nil {
		return err
	}
	return nil
}

//删除收藏夹中记录
func DeleteEntry(ctx context.Context, UserID, ColltID int64, targetEntry int64) error {
	res := make([]*Collection, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND id = ?", UserID, uint(ColltID)).Find(&res).First(&res).Error; err != nil {
		return err
	}
	if len(res) == 0 {
		return errors.New("no such Collection")
	}
	for i := 0; i < len(res[0].Entries); i++ {
		if res[0].Entries[i] == targetEntry {
			res[0].Entries = append(res[0].Entries[:i], res[0].Entries[i+1:]...)
			if err := DB.WithContext(ctx).Save(res).Error; err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("no such entry")
}

//设置收藏夹名
func SetName(ctx context.Context, UserID, ColltID int64, newName string) error {
	if err := DB.Model(&Collection{}).WithContext(ctx).Where("user_id = ? AND id = ?", UserID, uint(ColltID)).Update("name", newName).Error; err != nil {
		return err
	}
	return nil
}
