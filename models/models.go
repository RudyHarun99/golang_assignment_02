package models

import (
	"time"
)

type Item struct {
	ID					int
	Item_code		string	`gorm:"column:item_code; type:varchar(50)"`
	Description	string	`gorm:"column:description; type:varchar(255); default:null"`
	Quantity		int			`gorm:"column:quantity"`
	Order_id		int			`gorm:"column:order_id; type:varchar(50)"`
}

type Orders struct {
	ID 			 			int				`gorm:"type:autoIncrement"`
	Customer_name	string		`gorm:"column:customer_name; type:varchar(50)"`
	Ordered_at 		time.Time	`gorm:"column:created_at"`
	Items					[]Item		`gorm:"foreignKey:Order_id; references:ID"`
}