package main

import (
	"Aws/handlers"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Default will using on dev model
	r := gin.Default()

	//in busness model wll not print any thing ,just out put to the log file
	blog := r.Group("/blog")
	{
		blog.GET("", handlers.HandleMain)
		blog.GET("/about", handlers.HandleAbout)
	}

	r.Run(":8888")
}
