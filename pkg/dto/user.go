package dto

import "application/pkg/model"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"data"`
}

type RegisterRequest struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
	Email      string `json:"email"`
	Address    string `json:"address"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetUsersRequest struct {
	Q       *string         `query:"q"`
	Role    *model.UserRole `query:"role"`
	Page    uint            `query:"page"`
	PerPage uint            `query:"per_page"`
}

type GetUsersResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []model.User `json:"data"`
	Meta    model.Meta   `json:"meta"`
}

type UpdateUserRequest struct {
	Name       string         `json:"name"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Repassword string         `json:"repassword"`
	Role       model.UserRole `json:"role"`
	Email      string         `json:"email"`
	Address    string         `json:"address"`
}
