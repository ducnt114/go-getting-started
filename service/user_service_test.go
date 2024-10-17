package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-getting-started/dto"
	mockRepo "go-getting-started/mock/go-getting-started/repository"
	"go-getting-started/model"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func TestPasswordLogin(t *testing.T) {
	mockUserRepo := &mockRepo.MockUserRepository{}

	mockUserRepo.
		On("FindByName", mock.Anything, "user-1").
		Return(&model.User{Name: "user-1", Pass: "pass-1"}, nil)
	mockUserRepo.
		On("FindByName", mock.Anything, "user-2").
		Return(&model.User{Name: "user-2", Pass: "pass-2"}, nil)

	userService := &userServiceImpl{
		userRepo: mockUserRepo,
	}

	// valid
	resp, err := userService.PasswordLogin(
		context.Background(),
		"user-1",
		"pass-1")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal(t, resp.Meta.Code, http.StatusOK)

	// invalid pass
	_, err2 := userService.PasswordLogin(
		context.Background(),
		"user-2",
		"pass-2")
	if err2 == nil {
		t.Error(err2)
		t.Fail()
	}
}

func TestCreateUser(t *testing.T) {
	mockUserRepo := &mockRepo.MockUserRepository{}

	mockUserRepo.
		On("FindByName", mock.Anything, "user-existed").
		Return(&model.User{Name: "user-existed", Pass: "pass-1"}, nil)
	mockUserRepo.
		On("FindByName", mock.Anything, "user-not-existed").
		Return(nil, gorm.ErrRecordNotFound)
	mockUserRepo.
		On("FindByName", mock.Anything, "user-valid").
		Return(&model.User{Name: "user-valid", Pass: "pass-1"}, nil)

	userService := &userServiceImpl{
		userRepo: mockUserRepo,
	}

	// user existed
	_, err := userService.CreateUser(context.Background(), &dto.User{Name: "user-existed"})
	if err == nil {
		t.Fail()
	}

	// password length < 10
	_, err = userService.CreateUser(context.Background(),
		&dto.User{Name: "user-not-existed", Password: "123456789"})
	if err == nil {
		t.Fail()
	}

	// valid
	resp, err := userService.CreateUser(context.Background(),
		&dto.User{Name: "user-valid", Password: "1234567891011"})
	if err != nil {
		t.Fail()
	}
	if resp == nil {
		t.Fail()
	}
	assert.Equal(t, resp.Name, "user-valid")
}
