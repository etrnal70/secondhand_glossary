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

func TestAddTrait(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	traitRepo := NewTraitRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.Trait{})

	query := fmt.Sprintf("INSERT INTO `%s`", tableName)

	mock.ExpectBegin()
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	newTrait, err := traitRepo.AddTrait(model.Trait{Trait: "good battery", Context: "positive"})
	assert.NoError(t, err)
	assert.Equal(t, "good battery", newTrait.Trait)
	assert.Equal(t, "positive", newTrait.Context)
	assert.WithinDuration(t, time.Now(), newTrait.CreatedAt, 1*time.Second)
}

func TestDeleteTrait(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	traitRepo := NewTraitRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.Trait{})

	query := fmt.Sprintf("DELETE INTO `%s`", tableName)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(1).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()

	err := traitRepo.DeleteTrait(1)
	assert.NoError(t, err)
	assert.True(t, true)
}

func TestEditTrait(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	traitRepo := NewTraitRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.Trait{})

	query := fmt.Sprintf("UPDATE `%s`", tableName)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WillReturnResult(sqlmock.NewResult(0, 1)).
		WillReturnResult(driver.RowsAffected(1))
	mock.ExpectCommit()

	newTrait, err := traitRepo.EditTrait(model.Trait{ID: uint(1), Trait: "new_trait"})
	assert.NoError(t, err)
	assert.Equal(t, "new_trait", newTrait.Trait)
}

func TestGetTraitDevices(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	traitRepo := NewTraitRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.Trait{})

	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)

	rows := sqlmock.NewRows([]string{"id", "manufacturer", "lineup", "type"}).
		AddRow(1, "Lenovo", "Thinkpad", "T480").
		AddRow(2, "HP", "Elitebook", "840 G5")

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	devices, err := traitRepo.GetTraitDevices(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(devices))
}
