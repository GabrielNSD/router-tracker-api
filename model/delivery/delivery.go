package delivery

import (
	"api/model/user"
	"time"

	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	ID            string
	UserID        uint
	Client        user.User `gorm:"foreignKey:UserID"`
	StartingPoint string
	EndingPoint   string
	status        string
	OrderTime     *time.Time
	PickupTime    *time.Time
	DeliveryTime  *time.Time
}
