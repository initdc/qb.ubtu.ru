package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, _ := ioutil.ReadFile("sh/1key-docker-compose-ubuntu.sh")

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"sh": string(dat),
		})
	})

	router.HEAD("/", func(d *gin.Context) {
		d.File("sh/1key-docker-compose-ubuntu.sh")
	})

	router.Run(":80")
}
