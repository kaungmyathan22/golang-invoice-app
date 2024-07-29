package user

import (
	"github.com/kaungmyathan22/golang-invoice-app/app/lib"
)

type RegisterUserDTO struct {
	Username string `json:"username" binding:"required" valid:"required~Username is required"`
	Password string `json:"password" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
}

type LoginUserDTO struct {
	RegisterUserDTO
}

func (dto *RegisterUserDTO) ToModel() (*UserModel, error) {
	hashedPassword, err := lib.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	return &UserModel{Password: hashedPassword, Username: dto.Username}, nil
}
