package handler

import (
	"net/http"

	"github.com/faiz-muttaqin/lgs/backend/internal/database"
	"github.com/faiz-muttaqin/lgs/backend/internal/model"
	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []model.User
		database.DB.Preload("Roles").Find(&users)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Users fetched successfully",
			"data":    users,
		})
	}
}
