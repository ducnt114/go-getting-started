package dto

type User struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Bio   string  `json:"bio"`
	Books []*Book `json:"books"`
	Tag1  string  `json:"tag_1"`
}

type CreateUserReq struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Password string  `json:"password"`
	Books    []*Book `json:"books"`
}

type PasswordLoginResponse struct {
	Meta        *Meta  `json:"meta"`
	AccessToken string `json:"access_token"`
}

type ListUserResponse struct {
	Meta *Meta   `json:"meta"`
	Data []*User `json:"data"`
}
