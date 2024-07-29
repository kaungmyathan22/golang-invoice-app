package user_dto

import user_models "github.com/kaungmyathan22/golang-invoice-app/app/user/models"

type UserRegisterResponse struct {
	Username string `json:"username"`
}

func FromModel(model *user_models.UserModel) *UserRegisterResponse {
	return &UserRegisterResponse{
		Username: model.Username,
	}
}
