package repository

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepoDriver struct {
	DB *gorm.DB
}

func (d *userRepoDriver) GetUserProfile(userId uint) (user model.User, err error) {
	err = d.DB.Omit("Password").Find(&user, userId).Error
	return
}

func (d *userRepoDriver) GetUser(u model.UserLogin) (user model.User, err error) {
	err = d.DB.Where("email = ? AND  password = ?", u.Email, u.Password).Omit("Password").Take(&user).Error
	return
}

func (d *userRepoDriver) GetUsers() (users []model.User, err error) {
	err = d.DB.Omit("Password").Find(&users).Error
	return
}

func (d *userRepoDriver) RegisterUser(u model.UserRegister) (user model.User, err error) {
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password

	err = d.DB.Create(&user).Error
	return
}

func (d *userRepoDriver) UpdateUserProfile(u model.User) (user model.User, err error) {
	err = d.DB.Model(&user).Clauses(clause.Returning{}).Where("id = ?", u.ID).Updates(&u).Error
	return
}

func (d *userRepoDriver) DeleteUser(userId uint) (err error) {
	err = d.DB.Delete(&model.User{}, userId).Error
	return
}

func NewUserRepoDriver(db *gorm.DB) domain.UserRepository {
	return &userRepoDriver{
		DB: db,
	}
}
