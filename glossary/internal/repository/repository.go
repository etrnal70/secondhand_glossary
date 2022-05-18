package repository

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"
)

func NewGORMMock() (*sql.DB, sqlmock.Sqlmock) {
	dbConn, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error connecting to db stub: %s", err)
	}

	return dbConn, mock
}

func GetGORMTableName(db *gorm.DB, model interface{}) string {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	return stmt.Schema.Table
}
