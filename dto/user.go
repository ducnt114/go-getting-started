package dto

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PasswordLoginResponse struct {
	Meta        *Meta  `json:"meta"`
	AccessToken string `json:"access_token"`
}
