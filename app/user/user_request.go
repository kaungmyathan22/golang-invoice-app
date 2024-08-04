package user

import (
	"github.com/kaungmyathan22/golang-invoice-app/app/lib"
)

type RegisterUserDTO struct {
	Username string `json:"username" binding:"required" valid:"required~Username is required"`
	Email    string `json:"email" binding:"required" valid:"email,required~Valid email is required"`
	Password string `json:"password" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
}

func (dto *RegisterUserDTO) ToModel() (*UserModel, error) {
	hashedPassword, err := lib.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}
	return &UserModel{Password: hashedPassword, Username: dto.Username, Email: dto.Email}, nil
}

type LoginUserDTO struct {
	Email    string `json:"email" binding:"required" valid:"email,required~Valid email is required"`
	Password string `json:"password" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
}

type UpdateUserDTO struct {
	Username string `json:"username" binding:"required" valid:"required~Username is required"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"oldPassword" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
	NewPassword string `json:"newPassword" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
}

func (dto *ChangePasswordDTO) HashPassword() error {
	hashedPassword, err := lib.HashPassword(dto.NewPassword)
	if err != nil {
		return err
	}
	dto.NewPassword = hashedPassword
	return nil
}

type ForgotPasswordDTO struct {
	Email string `json:"email" binding:"required" valid:"email,required~Valid email is required"`
}
