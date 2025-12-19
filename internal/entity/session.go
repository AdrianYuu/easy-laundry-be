package entity

import "time"

type Session struct {
	Id          string     `gorm:"column:id;primaryKey;size:36"`
	LaundryName string     `gorm:"column:laundry_name;size:255;not null"`
	Status      Status     `gorm:"column:status;size:10;not null"`
	Note        *string    `gorm:"column:note;size:255;null"`
	DroppedAt   *time.Time `gorm:"column:dropped_at;null"`
	PickedUpAt  *time.Time `gorm:"column:picked_up_at;null"`
	UserId      string     `gorm:"column:user_id;size:36;not null"`
	User        User       `gorm:"foreignKey:UserId;references:Id"`
	Photos      []Photo    `gorm:"foreignKey:SessionId;references:Id"`
	CreatedAt   time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;not null"`
}

func (Session) TableName() string {
	return "sessions"
}
