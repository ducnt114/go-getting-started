package dto

type LoginResponse struct {
	Meta *Meta  `json:"meta"`
	Data *Token `json:"data"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type PasswordLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetTwoFaResponse struct {
	Meta *Meta              `json:"meta"`
	Data *TwoFaResponseData `json:"data"`
}

type TwoFaResponseData struct {
	Secret string `json:"secret"`
}

type SetupTwoFaRequest struct {
	OTP string `json:"otp"`
}

type SetupTwoFaResponse struct {
	Meta *Meta `json:"meta"`
}
