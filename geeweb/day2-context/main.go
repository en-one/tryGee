package main

import (
	"gee"
	"net/http"
)

// curl "http://localhost:8888/hello?name=geektutu"
// curl "http://localhost:8888/login" -X POST -d 'username=geektutu&password=1234'
// curl "http://localhost:8888/xxx"
func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>hello josiah</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s \n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8888")
}
