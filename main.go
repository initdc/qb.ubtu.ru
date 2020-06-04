package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	d, err := ioutil.ReadFile("sh/1key-docker-compose-ubuntu.sh")
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/index.html")

	router.GET("/", func(c *gin.Context) {
		ua := c.Request.Header.Get("User-Agent")
		if strings.Contains(ua, "like Gecko") {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"sh": string(d),
			})
		} else {
			c.File("sh/1key-docker-compose-ubuntu.sh")
		}
	})

	router.Run(":80")
}
