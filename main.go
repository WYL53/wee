package main

import (
	"net/http"

	"wee/wee"
)

func main() {
		r := wee.New()
		r.GET("/index", func(c *wee.Context) {
			c.HTML(http.StatusOK, "<h1>Index Page</h1>")
		})
		v1 := r.Group("/v1")
		{
			v1.GET("/", func(c *wee.Context) {
				c.HTML(http.StatusOK, "<h1>Hello wee</h1>")
			})
	
			v1.GET("/hello", func(c *wee.Context) {
				// expect /hello?name=geektutu
				c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
			})
		}
		v2 := r.Group("/v2")
		{
			v2.GET("/hello/:name", func(c *wee.Context) {
				// expect /hello/geektutu
				c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
			})
			v2.POST("/login", func(c *wee.Context) {
				c.JSON(http.StatusOK, wee.H{
					"username": c.PostForm("username"),
					"password": c.PostForm("password"),
				})
			})
	
		}
	
		r.Run(":9999")
}