package handlers

import (
	"Aws/Service"
	"github.com/gin-gonic/gin"
)

var App *Service.AppStaticsService

func init() {
	App = Service.NewAppStaticsService()
}

func HandleMain(c *gin.Context) {
	ip := c.ClientIP()
	o := gin.H{
		"ip": ip,
	}
	App.AppServ["blog"].InNumber++
	App.AppServ["blog"].OkOutNumber++
	c.JSON(200, o)
}

func HandleHistory(c *gin.Context) {
	ip := c.ClientIP()
	o := gin.H{
		"ip": ip,
	}
	App.AppServ["history"].InNumber++
	App.AppServ["history"].OkOutNumber++
	c.JSON(200, o)
}

func HandleWatcher(c *gin.Context) {
	c.JSON(200, gin.H{"errorcode": "0", "data": App.Format()})
}
