package repository

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"secondhand_glossary/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetUserProfile(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	userRepo := NewUserRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.User{})

	query := fmt.Sprintf("SELECT * FROM `%s` WHERE `%s`.`id` = ?", tableName, tableName)
	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "user_role", "created_at", "updated_at"}).
		AddRow(1, "Budi", "contact@budi.dev", "budi12345", "user", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	res, err := userRepo.GetUserProfile(1)
	assert.Equal(t, uint(1), res.ID)
	assert.Equal(t, "Budi", res.Username)
	assert.Equal(t, "budi12345", res.Password)
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	userRepo := NewUserRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.User{})

	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "user_role", "created_at", "updated_at"}).
		AddRow(1, "Adi", "contact@adi.dev", "adi12345", "user", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	user, err := userRepo.GetUser(model.UserLogin{Email: "contact@adi.dev", Password: "adi12345"})

	assert.NoError(t, err)
	assert.Equal(t, "contact@adi.dev", user.Email)
	assert.Equal(t, "adi12345", user.Password)
}

func TestGetUsers(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	userRepo := NewUserRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.User{})

	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "user_role", "created_at", "updated_at"}).
		AddRow(1, "Adi", "contact@adi.dev", "adi12345", "user", time.Now(), time.Now()).
		AddRow(2, "Budi", "contact@budi.dev", "budi12345", "user", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	users, err := userRepo.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
}

func TestRegisterUser(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	userRepo := NewUserRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.User{})

  query := fmt.Sprintf("INSERT INTO `%s`", tableName)

	mock.ExpectBegin()
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	newUser, err := userRepo.RegisterUser(model.UserRegister{Username: "Ado", Email: "contact@ado.dev", Password: "ado12345"})
	assert.NoError(t, err)
	assert.Equal(t, "contact@ado.dev", newUser.Email)
	assert.Equal(t, "ado12345", newUser.Password)
	assert.WithinDuration(t, time.Now(), newUser.CreatedAt, 1*time.Second)
}

func TestUpdateUserProfile(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	userRepo := NewUserRepoDriver(db)
	defer dbConn.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WillReturnResult(sqlmock.NewResult(0, 1)).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()

	newUser, err := userRepo.UpdateUserProfile(model.User{ID: uint(1), Username: "Ado", Email: "contact@ado.dev"})
	assert.NoError(t, err)
	assert.Equal(t, "Ado", newUser.Username)
	assert.Equal(t, "contact@ado.dev", newUser.Email)
}

func TestDeleteUser(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	userRepo := NewUserRepoDriver(db)
	defer dbConn.Close()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE").
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()

	err := userRepo.DeleteUser(1)
	assert.NoError(t, err)
	assert.True(t, true)
}
