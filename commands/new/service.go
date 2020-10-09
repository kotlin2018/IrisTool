package commands

var service = `package service

import (
	"{{.Appname}}/repositories"
)

type TestService struct {
	repo *repositories.TestRepositories
}

func NewTestService() *TestService {
	return &TestService{repo: repositories.NewTestRepositories()}
}`
