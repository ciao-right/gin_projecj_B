package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(page int, pageSize int, maps interface{}) (tags []Tag) {
	// Find 方法是用来查询符合条件的所有数据 Find(&tableName)
	//result := db.Find(&users)
	// SELECT * FROM users;
	//result.RowsAffected // returns found records count, equals `len(users)`
	//result.Error        // returns error

	// where 条件查询 map 实例见下
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
	db.Where(maps).Offset(page).Limit(pageSize).Find(&tags)
	return
}

// GetTagTotal 查询符合条件（maps）的tag的总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{CreatedBy: createdBy, Name: name, State: state})
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
func EditTag(id int, name string, state int, modifiedBy string) bool {
	var tag Tag
	db.Model(&tag).Where("ID = ?", id).Update(Tag{Name: name, State: state, ModifiedBy: modifiedBy})
	return true
}

//func EditTag(id int, data interface {}) bool {
//	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
//
//	return true
//}

func DeleteTag(id int) bool {
	var tag Tag
	db.Select("name").Where("id = ?", id).First(&tag)
	if tag.Name != "" {
		db.Delete(&tag, id)
		return true
	} else {
		panic("查无此记录")
		return false
	}
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
