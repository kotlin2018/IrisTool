package commands

var repositories = `package repositories

import (
	"{{.Appname}}/models"
	"github.com/jinzhu/gorm"
)

type TestRepositories struct {
	db *gorm.DB
}

func NewTestRepositories() *TestRepositories {
	return &TestRepositories{db: models.DB.Mysql}
}`
