package database

import (
	"errors"

	"github.com/desertp1geon/FlowDesk/internal/config"
	"github.com/desertp1geon/FlowDesk/internal/models/users"
	"github.com/desertp1geon/FlowDesk/pkg/results"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbContext struct {
	dbConfig config.DbConfig
	Context  *gorm.DB

	Users *gorm.DB
}

func InitDb(dbConfig config.DbConfig) results.Result[DbContext] {
	db, err := gorm.Open(postgres.Open(dbConfig.DbConnectionString), &gorm.Config{})
	if err != nil {
		return results.Error[DbContext](errors.New("Cannot init connection to database"))
	}
	return results.Ok(DbContext{
		dbConfig: dbConfig,
		Context:  db,

		Users: db.Model(&users.User{}),
	})
}

func (dc *DbContext) UpMigration() {
	err := dc.Context.AutoMigrate(
		&users.User{},
	)
	if err != nil {
		panic(err.Error())
	}
}
