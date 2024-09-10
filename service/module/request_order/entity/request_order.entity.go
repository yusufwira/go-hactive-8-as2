package entity

import "time"

type RequestOrder struct {
	OrderId      uint       `json:"order_id" gorm:"primaryKey"`
	CustomerName string     `json:"customer_name" gorm:"not null; type:varchar (255)"`
	OrderAt      *time.Time `json:"order_at" gorm:"not null`
	CreatedAt    time.Time  `json:"created_at" gorm:"not null`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"not null`
}

func (RequestOrder) TableName() string {
	return "public.request_orders" // Replace with your schema and table name
}

type Items struct {
	ItemCode    string    `json:"item_code" gorm:"not null; type: varchar (255)"`
	Description string    `json:"description" gorm:"not null; type: varchar (255)"`
	Quantity    int       `json:"quantity" gorm:"not null; type: float(8)"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null; type: timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null; type: timestamp"`
	ItemId      uint      `json:"id" gorm:"primaryKey"`
	OrderId     uint      `json:"order_id" gorm:"not null; type: integer"`
}

func (Items) TableName() string {
	return "public.items" // Replace with your schema and table name
}
