package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	ins, err := ioutil.ReadFile("sh/install.sh")
	rc, err := ioutil.ReadFile("sh/.zshrc")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/index.html")

	router.GET("/", func(c *gin.Context) {
		ua := c.Request.Header.Get("User-Agent")
		if strings.Contains(ua, "Gecko") {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"sh": string(ins),
			})
		} else {
			c.File("sh/install.sh")
		}
	})

	router.GET("/rc", func(d *gin.Context) {
		d.File("sh/.zshrc")
		fmt.Println(string(rc))
	})

	router.Run(":" + port)
}
