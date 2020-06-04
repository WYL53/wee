package main

import (
	"net/http"

	"wee/wee"
)

func main() {
	r := wee.New()
	r.GET("/", func(c *wee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello wee</h1>")
	})
	r.GET("/hello", func(c *wee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *wee.Context) {
		c.JSON(http.StatusOK, wee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("/hello/:name", func(c *wee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *wee.Context) {
		c.JSON(http.StatusOK, wee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}