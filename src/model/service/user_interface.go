package service

import (
	"github.com/lorraynefirme/api-golang/src/configuration/rest_err"
	"github.com/lorraynefirme/api-golang/src/model"
)

func NewUserDomainService() UserDomainService { return &userDomainService{} }

type userDomainService struct {}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface)	*rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}