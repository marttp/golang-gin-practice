package main

import (
	"gogin/httpd/handler"
	"gogin/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	// We can establish db here
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
