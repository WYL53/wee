package main

import (
	"log"
	"net/http"
	"time"

	"wee/wee"
)

func onlyForV2() wee.HandlerFunc {
	return func(c *wee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := wee.New()
	r.Use(wee.Logger()) // global midlleware
	r.GET("/", func(c *wee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello wee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *wee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}