package db

import (
	model "github.com/Salaton/screen-test.git/graph/model"
)

type Customer struct {
	ID          string   `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Phonenumber int      `json:"Phonenumber"`
	Email       string   `json:"Email"`
	Orders      []*Order `json:"orders" gorm:"foreignKey:CustomerID;"`
}

type Order struct {
	ID              string             `json:"id" gorm:"primaryKey"`
	Item            []*model.ItemInput `json:"item" gorm:"-"`
	Price           float64            `json:"price"`
	DateOrderPlaced string             `json:"date_order_placed"`
	CustomerID      string
}
