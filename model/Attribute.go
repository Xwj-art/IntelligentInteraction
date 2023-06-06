package model

import (
	"Ai/errmsg"
	"gorm.io/gorm"
)

type Attribute struct {
	gorm.Model
	Height uint `gorm:"type:int(5)" json:"height" label:"身高"`
	Weight uint `gorm:"type:int(5)" json:"weight" label:"体重"`
}

func CreateAttribute(data *Attribute) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERR
	}
	return errmsg.SUCCESS
}

func GetAttributeInfo(id int) (int, Attribute) {
	var attributeInfo Attribute
	err := db.Where("id = ?", id).First(&attributeInfo).Error
	if err != nil {
		return errmsg.ERR, attributeInfo
	}
	return errmsg.SUCCESS, attributeInfo
}

func DeleteAttribute(id int) int {
	err := db.Where("id = ?", id).Delete(&Attribute{}).Error
	if err != nil {
		return errmsg.ERR
	}
	return errmsg.SUCCESS
}

func UpdateAttribute(id int, data *Attribute) int {
	var maps = make(map[string]any)
	maps["height"] = data.Height
	maps["weight"] = data.Weight

	err := db.Model(&Attribute{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERR
	}
	return errmsg.SUCCESS
}
