package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
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

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("modified_on", time.Now().Unix())
	return nil
}

func ExistArticleById(id int) (isExist bool) {
	var article Article
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
	db.Model(&Article{}).Where(data).Count(&total)
	return total
}

func GetArticles(page int, pageSize int, maps interface{}) (list Article) {
	//preloads 因为list 是 Article 的实现，Article 里面有一个Tag 类型的字段，需要填充这个字段，所有使用preloads 填充在Tag字段内
	db.Preloads("Tag").Where(maps).Offset(pageSize).Limit(page).Find(&list)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ? ", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		// v.(t) 类型断言
		TagId:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}
