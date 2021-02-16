package user

import (
	"gorm.io/gorm"
)

// Repository interface
type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// NewRepository function
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
