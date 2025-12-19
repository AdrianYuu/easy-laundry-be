package entity

import "time"

type User struct {
	Id        string    `gorm:"column:id;primaryKey;size:36"`
	GoogleId  string    `gorm:"column:google_id;size:255;not null;unique"`
	Name      string    `gorm:"column:name;size:100;not null"`
	Sessions  []Session `gorm:"foreignKey:UserId;references:Id"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (User) TableName() string {
	return "users"
}
