package user_dto

type RegisterUserDTO struct {
	Username string `json:"username" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
	Password string `json:"password" binding:"required" valid:"required~Username is required"`
}

type LoginUserDTO struct {
	RegisterUserDTO
}
