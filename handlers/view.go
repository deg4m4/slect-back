package handlers

import (
	"net/http"
	"os"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/parthkax70/slect/utils"
)

func ViewFile(c *gin.Context) {

	fileName := c.Request.URL.Query().Get("file")
	token := c.Request.URL.Query().Get("token")

	username, err := utils.AuthToken(token)
	if err != nil {
		c.JSON(http.StatusNotFound, "file not found")
		return
	}

	user, err := user.Lookup(username)
	if err != nil {
		c.JSON(http.StatusNotFound, "file not found")
		return
	}

	file, err := os.Stat(user.HomeDir + "/" + fileName)

	if err != nil {
		c.JSON(http.StatusNotFound, "file not found")
		return
	}

	if file.IsDir() {
		c.JSON(http.StatusNotFound, "file not found")
		return
	}

	c.File(user.HomeDir + "/" + fileName)
}
