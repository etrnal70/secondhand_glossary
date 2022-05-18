package repository

import (
	"fmt"
	"regexp"
	"secondhand_glossary/internal/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetCategories(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	categoryRepo := NewCategoryRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.Category{})

	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)

	rows := sqlmock.NewRows([]string{"id", "category"}).
		AddRow(1, "laptop").
		AddRow(2, "smartphone")

	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WillReturnRows(rows)

	categories, err := categoryRepo.GetCategories()
	assert.NoError(t, err)
	assert.Equal(t, "laptop", categories[0].Category)
	assert.Equal(t, "smartphone", categories[1].Category)
	assert.Equal(t, 2, len(categories))
}

func TestGetCategoryDevices(t *testing.T) {
	dbConn, mock := NewGORMMock()
	db, _ := gorm.Open(mysql.Dialector{&mysql.Config{
		Conn:                      dbConn,
		SkipInitializeWithVersion: true,
	}})
	categoryRepo := NewCategoryRepoDriver(db)
	defer dbConn.Close()

	tableName := GetGORMTableName(db, &model.Device{})

	query := fmt.Sprintf("SELECT * FROM `%s`", tableName)
	rows := sqlmock.NewRows([]string{"id", "manufacturer", "lineup", "type"}).
		AddRow(1, "Lenovo", "Thinkpad", "T480").
		AddRow(2, "HP", "Elitebook", "840 G5")

  mock.ExpectQuery(regexp.QuoteMeta(query)).
  WillReturnRows(rows)

	devices, err := categoryRepo.GetCategoryDevices(1)
  assert.NoError(t, err)
  assert.Equal(t, 2, len(devices))
  assert.Equal(t, "Thinkpad", devices[0].Lineup)
  assert.Equal(t, "840 G5", devices[1].Type)
}
