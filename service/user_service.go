package service

import (
	"encoding/json"
	"go-getting-started/dto"
	"io"
	"net/http"
)

type UserService interface {
	GetUserByGender(gender string) *dto.UserResponse
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetUserByGender(gender string) *dto.UserResponse {
	username := fetchData("https://randomuser.me/api/?gender=" + gender)
	return &dto.UserResponse{
		Name: username,
	}
}

func fetchData(inputUrl string) string {
	resp, err := http.Get(inputUrl)
	if err != nil {
		return ""
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var userData UserData
	if err := json.Unmarshal(respBytes, &userData); err != nil {
		return ""
	}
	return userData.Results[0].Name.First
}

type UserData struct {
	Results []struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
	} `json:"results"`
}
