package models

type UserLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type UserRefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
