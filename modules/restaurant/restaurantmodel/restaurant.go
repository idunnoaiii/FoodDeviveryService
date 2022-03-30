package restaurantmodel

import "200lab/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name,omitempty" gorm:"column:name;"`
	Addr            string `json:"address,omitempty" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"id,omitempty" gorm:"column:id;"`
	Addr *string `json:"address,omitempty" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name,omitempty" gorm:"column:name;"`
	Addr string `json:"address,omitempty" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
