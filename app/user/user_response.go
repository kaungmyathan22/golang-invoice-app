package user

type UserRegisterResponse struct {
	Username string `json:"username"`
}

func FromModel(model *UserModel) *UserRegisterResponse {
	return &UserRegisterResponse{
		Username: model.Username,
	}
}

type UserLoginResponse struct {
	Token string     `json:"token"`
	User  UserEntity `json:"user"`
}
