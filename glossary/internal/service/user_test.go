package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO service function name should be the same as repo

func TestGetProfileDetails(t *testing.T) {
	expectedUser := model.User{ID: 1}

	mock := new(domain.MockUserRepository)
	mock.On("GetUserProfile", uint(1)).Return(expectedUser, nil)

	userService := NewUserService(mock)

	user, err := userService.GetProfileDetails(uint(1))

	assert.Nil(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
}

func TestRegister(t *testing.T) {
	expectedUser := model.User{Email: "test@user.com"}

	mock := new(domain.MockUserRepository)
	mock.On("RegisterUser", model.UserRegister{Email: "test@user.com"}).Return(expectedUser, nil)

	userService := NewUserService(mock)

	registeredUser, err := userService.Register(model.UserRegister{Email: "test@user.com"})
	assert.Nil(t, err)
	assert.Equal(t, expectedUser.Email, registeredUser.Email)
}

func TestUpdateProfile(t *testing.T) {
	expectedUser := model.User{Password: "newpassword"}

	mock := new(domain.MockUserRepository)
	mock.On("UpdateUserProfile", model.User{Password: "newpassword"}).Return(expectedUser, nil)

	traitService := NewUserService(mock)

	user, err := traitService.UpdateProfile(expectedUser)

	assert.Nil(t, err)
	assert.Equal(t, user.Password, expectedUser.Password)
}

func TestDeleteUser(t *testing.T) {
	mock := new(domain.MockUserRepository)
	mock.On("DeleteUser", uint(1)).Return(nil)

	traitService := NewUserService(mock)

	err := traitService.DeleteUser(1)

	assert.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	expectedUsers := []model.User{{ID: 1}, {ID: 2}}

	mock := new(domain.MockUserRepository)
	mock.On("GetUsers").Return(expectedUsers, nil)

	traitService := NewUserService(mock)

	users, err := traitService.GetUsers()

	assert.Nil(t, err)
	assert.Equal(t, users[0].ID, expectedUsers[0].ID)
	assert.Equal(t, users[1].ID, expectedUsers[1].ID)
}

func TestRegisterUser(t *testing.T) {
	expectedUser := model.User{Password: "newpassword"}

	mock := new(domain.MockUserRepository)
	mock.On("UpdateUserProfile", model.User{Password: "newpassword"}).Return(expectedUser, nil)

	traitService := NewUserService(mock)

	user, err := traitService.UpdateProfile(expectedUser)

	assert.Nil(t, err)
	assert.Equal(t, user.Password, expectedUser.Password)
}
