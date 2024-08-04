package user

import (
	"time"

	"gorm.io/gorm"
)

type UserEntity struct {
	ID                    uint           `json:"id"`
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt"`
	DeletedAt             gorm.DeletedAt `json:"deletedAt"`
	Username              string         `json:"username"`
	Email                 string         `json:"email"`
	Password              string         `json:"-"`
	LastLoggedInAt        time.Time      `json:"lastLoggedInAt"`
	LastPasswordUpdatedAt time.Time      `json:"lastPasswordUpdatedAt"`
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
