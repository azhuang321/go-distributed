package model

import "time"

type Category struct {
	ID               int       `json:"id" gorm:"column:id"`
	Name             string    `json:"name" gorm:"column:name"`
	ParentCategoryID int       `json:"parent_category_id" gorm:"column:parent_category_id"`
	Code             string    `json:"code" gorm:"column:code"`
	Desc             string    `json:"desc" gorm:"column:desc"`
	Level            int       `json:"level" gorm:"column:level"`
	IsTab            int8      `json:"is_tab" gorm:"column:is_tab"`
	AddTime          time.Time `json:"add_time" gorm:"column:add_time"`
}

func (m *Category) TableName() string {
	return "category"
}
