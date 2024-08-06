package user

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt"`
	LastLoggedInAt        time.Time      `json:"lastLoggedInAt"`
	LastPasswordUpdatedAt time.Time      `json:"lastPasswordUpdatedAt"`
	DeletedAt             gorm.DeletedAt `json:"deletedAt"`
	Username              string         `json:"username"`
	Email                 string         `json:"email"`
	Password              string         `json:"-"`
	ID                    uint           `json:"id"`
}

func UserEntityFromUserModel(model *UserModel) *UserEntity {
	return &UserEntity{
		ID:                    model.ID,
		CreatedAt:             model.CreatedAt,
		UpdatedAt:             model.UpdatedAt,
		DeletedAt:             model.DeletedAt,
		Username:              model.Username,
		Email:                 model.Email,
		Password:              model.Password,
		LastLoggedInAt:        model.LastLoggedInAt,
		LastPasswordUpdatedAt: model.LastPasswordUpdatedAt,
	}
}
