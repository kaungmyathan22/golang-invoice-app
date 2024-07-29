package user

type UserRegisterResponse struct {
	Username string `json:"username"`
}

func FromModel(model *UserModel) *UserRegisterResponse {
	return &UserRegisterResponse{
		Username: model.Username,
	}
}
