package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Articles struct {
	Model
	TagId      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (a Articles) BeforeCreate(scope gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (a Articles) BeforeUpdate(scope gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleById(id int) (isExist bool) {
	var article Articles
	db.Select("id").Where("id = ?", id).Find(&article)
	if article.ID > 0 {
		isExist = true
	} else {
		isExist = false
	}
	return
}

func GetArticlesTotal(data interface{}) int {
	var total int
	db.Model(&Articles{}).Where(data).Count(&total)
	return total
}
