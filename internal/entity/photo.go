package entity

import "time"

type Photo struct {
	Id        string    `gorm:"column:id;primaryKey;size:36"`
	Url       string    `gorm:"column:url;size:255;not null"`
	SessionId string    `gorm:"column:session_id;size:36;not null"`
	Session   Session   `gorm:"foreignKey:SessionId;references:Id"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (Photo) TableName() string {
	return "photos"
}
