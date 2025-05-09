package service

import (
	"fmt"

	"github.com/lorraynefirme/api-golang/src/configuration/logger"
	"github.com/lorraynefirme/api-golang/src/configuration/rest_err"
	"github.com/lorraynefirme/api-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	return nil
}