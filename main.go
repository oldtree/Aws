// Aws project main.go
package main

import (
	"github.com/oldtree/Aws/handlers"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Default will using on dev model
	r := gin.Default()

	blog := r.Group("/blog")
	{
		blog.GET("", handlers.HandleMain)
		blog.GET("/history", handlers.HandleHistory)
		blog.GET("/watcher", handlers.HandleWatcher)
	}
	browse := r.Group("/browse")
	{
		browse.GET("", handlers.HandleBrowse)
	}

	r.Run(":8888")
}
