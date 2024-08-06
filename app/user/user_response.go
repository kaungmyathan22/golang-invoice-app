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
	User  UserEntity `json:"user"`
	Token string     `json:"token"`
}
