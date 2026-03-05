package repository

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindById(db *gorm.DB, user *entity.User, id string) error {
	return db.Where("id = ?", id).Take(user).Error
}
