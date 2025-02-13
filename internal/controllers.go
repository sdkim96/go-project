package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

func GetMe(c *gin.Context) {
	user := User{
		ID:   uuid.New(),
		Name: "John Doe",
	}

	c.JSON(200, user)
}
