package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lorraynefirme/api-golang/src/configuration/rest_err"
)

func CreateUser(c *gin.Context){
	err := rest_err.NewBadRequestError("error creating user")
	c.JSON(err.Code, err)
}

