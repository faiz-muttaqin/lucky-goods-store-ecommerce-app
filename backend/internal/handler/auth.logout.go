package handler

import (
	"fmt"
	"net/http"

	"github.com/faiz-muttaqin/lgs/backend/internal/helper"
	"github.com/gin-gonic/gin"
)

func GetAuthLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		fmt.Println(userData)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Login successful",
		})
	}
}
