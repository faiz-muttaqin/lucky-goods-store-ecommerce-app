package websockets

import (
	"net/http"

	"github.com/faiz-muttaqin/lgs/backend/internal/database"
	"github.com/faiz-muttaqin/lgs/backend/internal/helper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func WebsocketHandlerGin(c *gin.Context) {
	userData, err := helper.GetFirebaseUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	logrus.Printf("WebSocket connection attempt from user: %s (ID: %d)", userData.Email, userData.ID)
	HandleWebSocket(c.Writer, c.Request, userData, database.DB)
}
