package dto

type UserResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  string `json:"age"`
}
