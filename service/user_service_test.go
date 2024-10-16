package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPasswordLogin(t *testing.T) {
	userService := &userServiceImpl{}
	resp, err := userService.PasswordLogin(
		context.Background(),
		"user",
		"pass")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, resp.Meta.Code, http.StatusOK)
}
