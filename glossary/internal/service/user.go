package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Repo domain.UserRepository
}

func (s *userService) GetProfileDetails(userId uint) (user model.User, err error) {
	user, err = s.Repo.GetUserProfile(userId)
	return
}

func (s *userService) Login(l model.UserLogin) (user model.User, err error) {
  // TODO bcrypt.CompareHashAndPassword(hashedPassword []byte, password []byte)
	panic("unimplemented")
}

func (s *userService) Logout(userId uint) (err error) {
	panic("unimplemented")
}

func (s *userService) Register(r model.UserRegister) (user model.User, err error) {
	// Encrypt password with salt
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	r.Password = string(hashedPass)

	user, err = s.Repo.RegisterUser(r)
	return
}

func (s *userService) UpdateProfile(u model.User) (user model.User, err error) {
	user, err = s.Repo.UpdateUserProfile(u)
	return
}

func (s *userService) DeleteUser(userId uint) (err error) {
	err = s.Repo.DeleteUser(userId)
	return
}

func (s *userService) GetUser(u model.UserLogin) (user model.User, err error) {
	user, err = s.Repo.GetUser(u)
	return
}

func (s *userService) GetUserProfile(userId uint) (user model.User, err error) {
	user, err = s.Repo.GetUserProfile(userId)
	return
}

func (s *userService) GetUsers() (users []model.User, err error) {
	users, err = s.Repo.GetUsers()
	return
}

func (s *userService) RegisterUser(u model.UserRegister) (user model.User, err error) {
	user, err = s.Repo.RegisterUser(u)
	return
}

func (s *userService) UpdateUserProfile(u model.User) (user model.User, err error) {
	user, err = s.Repo.UpdateUserProfile(u)
	return
}

func NewUserService(r domain.UserRepository) domain.UserService {
	return &userService{
		Repo: r,
	}
}
