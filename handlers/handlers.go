package handlers

import (
	"github.com/gin-gonic/gin"
)

func HandleMain(c *gin.Context) {
	ip := c.ClientIP()
	o := gin.H{
		"ip": ip,
	}
	c.JSON(200, o)
}

func HandleAbout(c *gin.Context) {
	ip := c.ClientIP()
	o := gin.H{
		"ip": ip,
	}
	c.JSON(200, o)
}
