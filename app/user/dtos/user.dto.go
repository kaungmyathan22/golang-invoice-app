package user_dto

type RegisterUserDTO struct {
	Username string `json:"username" binding:"required" valid:"required~Username is required"`
	Password string `json:"password" binding:"required" valid:"sixToEightDigitAlphanumericPasswordValidator~Password must be between 6 to 8 alphanumeric characters"`
}

type LoginUserDTO struct {
	RegisterUserDTO
}
