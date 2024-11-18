package user_signup

import (
	"context"
	"fmt"
	"log"
)

func SaveUserToDatabase(ctx context.Context, req UserSignupRequest) (string, error) {
	log.Printf("save user to db for user name: %v\n\n", req.Username)
	msg := fmt.Sprintf("SaveUserToDatabase success with name: %v, email: %v", req.Username, req.Email)
	return msg, nil
}

func SendWelcomeEmail(ctx context.Context, req UserSignupRequest) (string, error) {
	if req.Email == "error@gmail.com" {
		return "", fmt.Errorf("error when send welcome email for email: %v", req.Email)
	}

	log.Printf("send welcome email for user name: %v\n\n", req.Username)
	msg := fmt.Sprintf("SendWelcomEmail to email: %v", req.Email)
	return msg, nil
}
