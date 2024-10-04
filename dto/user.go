package dto

type UserResponse struct {
	ID   string `json:"id,omitempty" validate:"required,gt=10"`
	Name string `json:"name"`
	Age  string `json:"age"`
}
