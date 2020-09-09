package model

import "github.com/jinzhu/gorm"

type Article struct {
	Category Category
	gorm.Model
	Title   string `gorm:"type:varchar(100); not null" json:"title"`
	Cid     string `gorm:"type:int; not null" json:"cid"`
	Desc    string `gorm:"type:varchar(50);" json:"desc"`
	Content string `gorm:"type:longtext; not null" json:"content"`
	Image   string `gorm:"type:varchar(200)" json:"image"`
}
