package controllers

import (
	"auth-example/models"
	"auth-example/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRandomUser(c *gin.Context) {
	userRetrieved, err := services.GetRandomUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := &models.RandomUser{
		Gender: userRetrieved.Results[0].Gender,
		Name:   userRetrieved.Results[0].Name.First + " " + userRetrieved.Results[0].Name.Last,
		Email:  userRetrieved.Results[0].Email,
	}

	c.JSON(http.StatusOK, &models.RandomUserResponse{
		Results: []models.RandomUser{*user},
	})
}
