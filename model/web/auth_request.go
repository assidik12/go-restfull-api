package web

type AuthLoginRequest struct {
	Email    string `validate:"required,max=200,min=1" json:"email"`
	Password string `validate:"required,min=1" json:"password"`
}

type AuthRegisterRequest struct {
	Username string `validate:"required,max=200,min=1" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required,min=1" json:"password"`
}

type AuthUpdateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}

type AuthRegisterResponse struct {
	Message string `json:"message"`
}

type AuthUpdateResponse struct {
	Message string `json:"message"`
}
