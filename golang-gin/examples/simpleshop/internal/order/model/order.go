package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrgsrylm/simpleshop/pkg/utils"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "new"
	OrderStatusInProgress OrderStatus = "in-progress"
	OrderStatusDone       OrderStatus = "done"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID         string       `json:"id" gorm:"unique;not null;index;primary_key"`
	Code       string       `json:"code"`
	TotalPrice float64      `json:"total_price"`
	Status     OrderStatus  `json:"status"`
	UserID     string       `json:"user_id"`
	User       *User
	Lines      []*OrderLine `json:"lines"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  *time.Time   `json:"deleted_at" gorm:"index"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New().String()
	order.Code = utils.GenerateCode("SO")

	if order.Status == "" {
		order.Status = OrderStatusNew
	}

	return nil
}
